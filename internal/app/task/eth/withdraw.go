package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	ent_withdraw "github.com/skyisboss/pay-system/ent/withdraw"

	// "github.com/skyisboss/pay-system/internal/app/kms"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

// 检测用户提币
// 从withdraw表获取提币数据，处理后添加签名到transfer表
func (p *Provider) CheckWithdraw() {
	log := p.ioc.Logger()
	ctx := context.Background()
	if ok := p.ioc.ApprunService().Lock(ctx, "CheckWithdraw"); !ok {
		return
	}
	defer p.ioc.ApprunService().UnLock(ctx, "CheckWithdraw")

	// 获取币种配置信息
	cfgs, err := p.ioc.BlockchainService().GetByChain(ctx, wallet.ETH)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}
	chainIDs := util.MapSlice(cfgs, func(e *ent.Blockchain) uint64 {
		return e.ID
	})
	chainMap := util.KeyFunc(cfgs, func(e *ent.Blockchain) ChainID {
		return ChainID(e.ID)
	})

	// 获取待提币数据
	withdrawRows, err := p.ioc.WithdrawService().ListWithdrawByHandleStatus(ctx, chainIDs)
	if err != nil {
		log.Error().Err(err).Msg("ListByHandleStatus")
		return
	}
	if len(withdrawRows) <= 0 {
		// 没有要处理的提币
		log.Info().Err(err).Msg("暂无处理的数据")
		return
	}

	// 热钱包信息
	hotWalletMap := make(map[Address]*Wallet)
	balanceMap := make(map[ChainID]*big.Int)
	for _, cid := range chainIDs {
		vcid := cid
		cfg := chainMap[ChainID(vcid)]
		htAddress := Address(cfg.HotAddress)
		info, err := p.GetWallet(ctx, []Address{htAddress})
		if err != nil {
			log.Error().Err(err).Msg("GetWallet")
			return
		}
		// 获取nonce值
		nonce, err := getAddressNonce(ctx, p.ioc, htAddress)
		if nil != err {
			p.ioc.Logger().Error().Err(err).Msg("getAddressNonce")
			return
		}
		// 获取余额
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			balance, err := p.RpcBalanceAt(ctx, cfg.HotAddress)
			if err != nil {
				p.ioc.Logger().Error().Err(err).Msg("RpcBalanceAt")
				return
			}
			balanceMap[ChainID(vcid)] = balance
		} else {
			balance, err := p.RpcTokenBalanceAt(ctx, cfg.TokenAbi, cfg.TokenAddress, cfg.HotAddress)
			if err != nil {
				p.ioc.Logger().Error().Err(err).Msg("RpcBalanceAt")
				return
			}
			balanceMap[ChainID(vcid)] = balance
		}

		hotWalletMap[htAddress] = &Wallet{
			Address:    htAddress,
			PrivateKey: info[htAddress].Private,
			Nonce:      nonce,
			Mu:         sync.RWMutex{},
			BalanceMap: balanceMap,
		}
	}

	handler := &Withdraw{
		ioc:       p.ioc,
		ctx:       ctx,
		cfgMap:    chainMap,
		hotWallet: hotWalletMap,
	}
	handler.HandleWithdraw(withdrawRows)

}

// 热钱包
type Wallet struct {
	Address    Address
	PrivateKey *ecdsa.PrivateKey
	Nonce      int64
	BalanceMap map[ChainID]*big.Int
	Mu         sync.RWMutex
}

type Withdraw struct {
	ctx    context.Context
	ioc    *ioc.Container
	cfgMap map[ChainID]*ent.Blockchain
	// 热钱包用作交易签名，作为from地址，申请提币从这里扣除
	hotWallet map[Address]*Wallet
}

type CountInfo struct {
	// 数据id
	IDs []uint64
	// 所属网络
	Chain ChainID
	// 发送总金额 wei单位
	SendAmount *big.Int
}

func (h *Withdraw) HandleWithdraw(rows []*ent.Withdraw) {
	// 映射相同网络币种的提币信息
	withdrawChainMap := make(map[ChainID][]*ent.Withdraw)
	for _, itemRow := range rows {
		withdrawChainMap[ChainID(itemRow.ChainID)] = append(withdrawChainMap[ChainID(itemRow.ChainID)], itemRow)
	}

	// 归并相同币种和to地址
	for chain, withdrawRows := range withdrawChainMap {
		var toRows []Address
		toMap := make(map[Address]*CountInfo)
		rows = withdrawRows
		for _, itemRow := range rows {
			info := toMap[Address(itemRow.ToAddress)]
			if info == nil {
				info = &CountInfo{
					IDs:        []uint64{},
					Chain:      ChainID(itemRow.ChainID),
					SendAmount: new(big.Int),
				}
				toMap[Address(itemRow.ToAddress)] = info
			}
			info.IDs = append(info.IDs, itemRow.ID)
			info.SendAmount.Add(info.SendAmount, itemRow.AmountRaw.BigInt())
			if !util.InArray(toRows, Address(itemRow.ToAddress)) {
				toRows = append(toRows, Address(itemRow.ToAddress))
			}
		}
		vchain := chain
		cfg := h.cfgMap[vchain]
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			// 主币
			h.WithdrawCoin(toRows, toMap)
		} else {
			// 代币
			h.WithdrawToken(toRows, toMap)
		}
	}
}

func (h *Withdraw) WithdrawCoin(toRows []Address, toMap map[Address]*CountInfo) {
	for to, info := range toMap {
		cfg := h.cfgMap[info.Chain]
		// 发送金额
		sendAmount := info.SendAmount
		// 热钱包
		hotWallet := h.hotWallet[Address(cfg.HotAddress)]
		hotWallet.Mu.Lock()
		private := hotWallet.PrivateKey
		nonce := hotWallet.Nonce
		h.hotWallet[Address(cfg.HotAddress)].Nonce += 1
		hotWallet.Mu.Unlock()

		gas := cfg.GasPrice
		gasLimit := int64(21000)
		gasPrice := gas.UserGasPrice
		tipPrice := big.NewInt(gas.TipFeePrice)
		NetworkID, err := h.ioc.EthClient().NetworkID(h.ctx)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("err")
			return
		}

		h.NewSignTransaction(info, &NewTransaction{
			nonce:     nonce,
			to:        to.String(),
			amount:    sendAmount,
			gasLimit:  gasLimit,
			gasFeeCap: gasPrice,
			gasTipCap: tipPrice.Int64(),
			chainID:   NetworkID.Int64(),
			prv:       private,
		})
	}
}

func (h *Withdraw) WithdrawToken(toRows []Address, toMap map[Address]*CountInfo) {
	for to, info := range toMap {
		cfg := h.cfgMap[info.Chain]
		// 发送金额
		sendAmount := info.SendAmount
		// 热钱包
		hotWallet := h.hotWallet[Address(cfg.HotAddress)]
		hotWallet.Mu.Lock()
		private := hotWallet.PrivateKey
		nonce := hotWallet.Nonce
		h.hotWallet[Address(cfg.HotAddress)].Nonce += 1
		hotWallet.Mu.Unlock()

		gas := cfg.GasPrice
		gasLimit := int64(2100000)
		gasPrice := gas.UserGasPrice
		tipPrice := big.NewInt(gas.TipFeePrice)
		NetworkID, _ := h.ioc.EthClient().NetworkID(h.ctx)

		contractAbi, err := abi.JSON(strings.NewReader(cfg.TokenAbi))
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("读取abi")
			return
		}
		input, err := contractAbi.Pack(
			"transfer",
			to.HexAddress(),
			sendAmount,
		)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("contractAbi.Pack")
			return
		}

		err = h.NewSignTransaction(info, &NewTransaction{
			nonce:     nonce,
			to:        to.String(),
			amount:    sendAmount,
			gasLimit:  gasLimit,
			gasFeeCap: gasPrice,
			gasTipCap: tipPrice.Int64(),
			chainID:   NetworkID.Int64(),
			prv:       private,
			data:      input,
		})
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("NewSignTransaction")
			return
		}
	}
}

// 生成交易信息
func (h *Withdraw) NewSignTransaction(info *CountInfo, newTx *NewTransaction) error {
	cfg := h.cfgMap[info.Chain]
	sendAmount := newTx.amount
	sendFrom := cfg.HotAddress
	sendTo := newTx.to
	// 代币交易
	if newTx.data != nil {
		sendAmount = newTx.amount
		// 发送到合约地址
		newTx.to = cfg.TokenAddress
		//  token转账时金额为0，因为代币写在input数据里
		newTx.amount = big.NewInt(0)
	}
	// 创建交易
	signedTx, err := NewSignTransaction(newTx)
	if err != nil {
		h.ioc.Logger().Error().Err(err).Msg("NewSignTransaction")
		return err
	}

	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		h.ioc.Logger().Error().Err(err).Msg("MarshalBinary")
		return err
	}
	txHex := hex.EncodeToString(rawTxBytes)
	txHash := strings.ToLower(signedTx.Hash().Hex())
	amountRaw, _ := util.BigintToDecimal(sendAmount)
	amountStr, _ := util.WeiToEth(sendAmount, int32(cfg.Decimals))

	// 创建存入Transfer的数据
	var sendRows []*ent.Transfer
	sendHashMap := make(map[uint64]string)
	for rowIndex, rowID := range info.IDs {
		// 只有第一条数据需要发送，其余数据为占位数据
		if rowIndex == 0 {
			sendRows = append(sendRows, &ent.Transfer{
				ChainID:      cfg.ID,
				RelatedType:  transfer.RelatedTypeWithdraw,
				RelatedID:    rowID,
				TxID:         txHash,
				FromAddress:  sendFrom,
				ToAddress:    sendTo,
				AmountStr:    amountStr,
				AmountRaw:    amountRaw,
				Gas:          newTx.gasLimit,
				GasPrice:     newTx.gasFeeCap,
				Nonce:        newTx.nonce,
				Hex:          txHex,
				HandleStatus: transfer.HandleStatusInit.Int64(),
				HandleMsg:    transfer.HandleStatusInit.String(),
				HandleTime:   time.Now(),
			})
		} else {
			// 占位数据
			sendRows = append(sendRows, &ent.Transfer{
				ChainID:      cfg.ID,
				RelatedType:  transfer.RelatedTypeWithdraw,
				RelatedID:    rowID,
				TxID:         txHash,
				FromAddress:  sendFrom,
				ToAddress:    sendTo,
				AmountStr:    "0",
				AmountRaw:    decimal.NewFromInt(0),
				Gas:          0,
				GasPrice:     0,
				Nonce:        -1,
				Hex:          "",
				HandleStatus: transfer.HandleStatusInit.Int64(),
				HandleMsg:    transfer.HandleStatusInit.String(),
				HandleTime:   time.Now(),
			})
		}
		sendHashMap[rowID] = txHash
	}
	// 开启事务
	tx, err := h.ioc.DBClient().Tx(h.ctx)
	if err != nil {
		h.ioc.Logger().Error().Err(err).Msg("new transactional client")
		return err
	}
	// 事务回滚
	rollback := func(tx *ent.Tx, err error) error {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return err
	}

	for id, txHash := range sendHashMap {
		// 更新提款状态
		_, err := tx.Withdraw.Update().
			Where(ent_withdraw.IDEQ(id)).
			SetTxHash(txHash).
			SetHandleMsg(withdraw.WithdrawStatusHex.String()).
			SetHandleStatus(withdraw.WithdrawStatusHex.Int64()).
			SetHandleTime(time.Now()).
			Save(h.ctx)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("err")
			return rollback(tx, err)
		}
	}

	// 插入待发送数据
	rows := []*ent.TransferCreate{}
	for _, v := range sendRows {
		row := tx.Transfer.Create().
			SetRelatedID(v.RelatedID).
			SetRelatedType(v.RelatedType).
			SetTxID(v.TxID).
			SetChainID(v.ChainID).
			SetFromAddress(v.FromAddress).
			SetToAddress(v.ToAddress).
			SetAmountStr(v.AmountStr).
			SetAmountRaw(v.AmountRaw).
			SetHandleStatus(v.HandleStatus).
			SetHandleMsg(v.HandleMsg).
			SetHandleTime(time.Now()).
			SetCreatedAt(time.Now()).
			SetGas(v.Gas).
			SetGasPrice(v.GasPrice).
			SetNonce(v.Nonce).
			SetHex(v.Hex)

		rows = append(rows, row)
	}
	_, err = tx.Transfer.CreateBulk(rows...).Save(h.ctx)
	if err != nil {
		h.ioc.Logger().Error().Err(err).Msg("err")
		return rollback(tx, err)
	}

	// 提交事物
	return tx.Commit()
}
