package tron

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	ent_withdraw "github.com/skyisboss/pay-system/ent/withdraw"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/rpc/tronrpc"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
	"google.golang.org/protobuf/proto"
)

// 检测用户提币
// 从withdraw表获取提币数据，处理后添加签名到transfer表
// TODO::1、最小提币金额验证 2、热钱包余额是否足够构建交易
func (p *Provider) CheckWithdraw() {
	log := p.Ioc().Logger()
	ctx := context.Background()
	if ok := p.Ioc().ApprunService().Lock(ctx, "CheckWithdraw-tron"); !ok {
		return
	}
	defer p.Ioc().ApprunService().UnLock(ctx, "CheckWithdraw-tron")

	// 获取币种配置信息
	cfgs, err := p.Ioc().BlockchainService().GetByChain(ctx, wallet.TRON)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}
	chainIDs := util.MapSlice(cfgs, func(e *ent.Blockchain) uint64 {
		return e.ID
	})
	chainMap := util.KeyFunc(cfgs, func(e *ent.Blockchain) wallet.ChainID {
		return wallet.ChainID(e.ID)
	})

	// 获取待提币数据
	withdrawRows, err := p.Ioc().WithdrawService().ListWithdrawByHandleStatus(ctx, chainIDs)
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
	hotWalletMap := make(map[wallet.Address]*HotWallet)
	balanceMap := make(map[wallet.ChainID]*big.Int)
	for _, cid := range chainIDs {
		vcid := cid
		cfg := chainMap[wallet.ChainID(vcid)]
		htAddress := wallet.Address(cfg.HotAddress)
		dbAddress, err := p.Ioc().AddressService().GetByAddress(ctx, cfg.HotAddress)
		if err != nil {
			log.Error().Err(err).Msg("GetByAddress")
			continue
		}
		info := &wallet.Wallet{
			Address:    cfg.HotAddress,
			PrivateKey: dbAddress.Password,
			Blockchain: wallet.TRON,
			UUID:       dbAddress.UUID,
		}
		PrivateKey, err := info.DecodePrivateKey(p.Ioc().Config().Providers.SaltKey)
		if err != nil || PrivateKey == "" {
			log.Error().Err(err).Msg("EncodePrivateKey")
			continue
		}
		info.PrivateKey = PrivateKey

		// 获取余额
		Provider := p.Ioc().WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			account, err := Provider.Client.GetTrxBalance(cfg.HotAddress)
			if err != nil {
				p.Ioc().Logger().Error().Err(err).Msg("RpcBalanceAt")
				return
			}
			balanceMap[wallet.ChainID(vcid)] = new(big.Int).SetInt64(account.Balance)
		} else {
			balance, err := Provider.Client.GetTrc20Balance(cfg.HotAddress, cfg.TokenAddress)
			if err != nil {
				p.Ioc().Logger().Error().Err(err).Msg("GetTrc20Balance")
				return
			}
			balanceMap[wallet.ChainID(vcid)] = balance
		}

		hotWalletMap[htAddress] = &HotWallet{
			Address:    htAddress,
			PrivateKey: PrivateKey,
			BalanceMap: balanceMap,
			Nonce:      int64(20 * 3600 * 1000), //延长20小时
		}
	}

	handler := &Withdraw{
		ioc:       p.Ioc(),
		ctx:       ctx,
		cfgMap:    chainMap,
		hotWallet: hotWalletMap,
	}
	handler.HandleWithdraw(withdrawRows)
}

// 热钱包
type HotWallet struct {
	Address    wallet.Address
	PrivateKey string
	Nonce      int64
	BalanceMap map[wallet.ChainID]*big.Int
}

type Withdraw struct {
	ctx    context.Context
	ioc    *ioc.Container
	cfgMap map[wallet.ChainID]*ent.Blockchain
	// 热钱包用作交易签名，作为from地址，申请提币从这里扣除
	hotWallet map[wallet.Address]*HotWallet
}

type CountInfo struct {
	// 数据id
	IDs []uint64
	// 所属网络
	Chain wallet.ChainID
	// 发送总金额 wei单位
	SendAmount *big.Int
}

func (h *Withdraw) HandleWithdraw(rows []*ent.Withdraw) {
	// TronProvider := h.ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)

	/*对相同chain的相同提币地址做归并处理*/

	// 1、先提取相同网络
	withdrawMap := make(map[wallet.ChainID][]*ent.Withdraw)
	for _, itemRow := range rows {
		withdrawMap[wallet.ChainID(itemRow.ChainID)] = append(withdrawMap[wallet.ChainID(itemRow.ChainID)], itemRow)
	}
	// 2、再做归并
	for chain, withdrawRows := range withdrawMap {
		// chainID := chain
		cfg := h.cfgMap[chain]

		if cfg.Status != 1 {
			h.ioc.Logger().Warn().
				Str("chain-type", cfg.Types).
				Str("chain-id", wallet.ChainID(cfg.ID).String()).
				Msg("系统关闭检测")
			continue
		}

		// var toRows []wallet.Address
		toAddressMap := make(map[wallet.Address]*CountInfo)
		for _, itemRow := range withdrawRows {
			info := toAddressMap[wallet.Address(itemRow.ToAddress)]
			if info == nil {
				info = &CountInfo{
					IDs:        []uint64{},
					Chain:      wallet.ChainID(itemRow.ChainID),
					SendAmount: new(big.Int),
				}
				toAddressMap[wallet.Address(itemRow.ToAddress)] = info
			}
			info.IDs = append(info.IDs, itemRow.ID)
			info.SendAmount.Add(info.SendAmount, itemRow.AmountRaw.BigInt())
			// if !util.InArray(toRows, wallet.Address(itemRow.ToAddress)) {
			// 	toRows = append(toRows, wallet.Address(itemRow.ToAddress))
			// }
		}

		// 创建存入Transfer的数据
		/*
			// 热钱包
			hotWallet := h.hotWallet[wallet.Address(cfg.HotAddress)]
			fromAddress := hotWallet.Address.String()
			privateKey := hotWallet.PrivateKey
			nonce := hotWallet.Nonce
			var sendRows []*ent.Transfer
			sendHashMap := make(map[uint64]string)
			// 处理交易
			for to, info := range toAddressMap {
				toAddress := to.String()
				amountInt := info.SendAmount.Int64()
				amountBigInt := new(big.Int).SetInt64(amountInt)
				amountRaw := decimal.NewFromInt(amountInt)
				amountStr, err := util.WeiToEth(amountBigInt, int32(cfg.Decimals))
				if err != nil {
					h.ioc.Logger().Error().Err(err).Msg("WeiToEth")
					continue
				}
				// 创建交易数据
				var tx *api.TransactionExtention
				if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
					tx, err = TronProvider.Client.GRPC.Transfer(fromAddress, toAddress, amountInt)
					if err != nil {
						h.ioc.Logger().Error().Err(err).Msg("Transfer")
						continue
					}
				} else {
					tx, err = TronProvider.Client.GRPC.TRC20Send(fromAddress, toAddress, cfg.TokenAddress, amountBigInt, 3000000000)
					if err != nil {
						h.ioc.Logger().Error().Err(err).Msg("TRC20Send")
						continue
					}
				}
				// 延长交易过期时间  https://developers.tron.network/docs/tron-protocol-transaction
				// raw_data.expiration - 交易过期时间，超过该时间交易将不再被打包。 如果交易是通过调用java-tron API创建的，
				// 则其过期时间将由节点自动设置为该节点最新块的时间戳加上60秒的值。 过期时间间隔可以在节点的配置文件中修改，最大值不能超过24小时。
				tx.Transaction.RawData.Expiration += nonce
				// txHash 一直出错的原因 https://github.com/fbsobreira/gotron-sdk/issues/93
				TronProvider.Client.GRPC.UpdateHash(tx)

				// 对数据签名
				signTx, err := tronrpc.SignTransaction(tx.Transaction, privateKey)
				if err != nil {
					h.ioc.Logger().Error().Err(err).Msg("SignTransaction")
					continue
				}
				// 转换原始数据
				rawData, err := proto.Marshal(signTx)
				if err != nil {
					h.ioc.Logger().Error().Err(err).Msg("Marshal")
					continue
				}
				txHex := hex.EncodeToString(rawData)
				txHash := common.Bytes2Hex(tx.GetTxid())

				for rowIndex, rowID := range info.IDs {
					sendHashMap[rowID] = txHash
					// 只有第一条数据需要发送，其余数据为占位数据
					if rowIndex == 0 {
						sendRows = append(sendRows, &ent.Transfer{
							ChainID:      cfg.ID,
							RelatedType:  transfer.RelatedTypeWithdraw,
							RelatedID:    rowID,
							TxID:         txHash,
							FromAddress:  fromAddress,
							ToAddress:    toAddress,
							AmountStr:    amountStr,
							AmountRaw:    amountRaw,
							Gas:          0,
							GasPrice:     0,
							Nonce:        nonce,
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
							FromAddress:  fromAddress,
							ToAddress:    toAddress,
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
				}
			}
		*/

		sendRows, sendHashMap := h.SignTransaction(toAddressMap, cfg)

		// 添加数据
		err := h.AddWithdraw(sendRows, sendHashMap)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("AddWithdraw")
		}
	}
}

// 处理交易签名
func (h *Withdraw) SignTransaction(toAddressMap map[wallet.Address]*CountInfo, cfg *ent.Blockchain) ([]*ent.Transfer, map[uint64]string) {
	TronProvider := h.ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)
	// 热钱包
	hotWallet := h.hotWallet[wallet.Address(cfg.HotAddress)]
	fromAddress := hotWallet.Address.String()
	privateKey := hotWallet.PrivateKey
	nonce := hotWallet.Nonce
	// 创建存入Transfer的数据
	var sendRows []*ent.Transfer
	sendHashMap := make(map[uint64]string)

	for to, info := range toAddressMap {
		toAddress := to.String()
		amountInt := info.SendAmount.Int64()
		amountBigInt := new(big.Int).SetInt64(amountInt)
		amountRaw := decimal.NewFromInt(amountInt)
		amountStr, err := util.WeiToEth(amountBigInt, int32(cfg.Decimals))
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("WeiToEth")
			continue
		}
		// 创建交易数据
		var tx *api.TransactionExtention
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			tx, err = TronProvider.Client.GRPC.Transfer(fromAddress, toAddress, amountInt)
			if err != nil {
				h.ioc.Logger().Error().Err(err).Msg("Transfer")
				continue
			}
		} else {
			tx, err = TronProvider.Client.GRPC.TRC20Send(fromAddress, toAddress, cfg.TokenAddress, amountBigInt, 3000000000)
			if err != nil {
				h.ioc.Logger().Error().Err(err).Msg("TRC20Send")
				continue
			}
		}
		// 延长交易过期时间  https://developers.tron.network/docs/tron-protocol-transaction
		// raw_data.expiration - 交易过期时间，超过该时间交易将不再被打包。 如果交易是通过调用java-tron API创建的，
		// 则其过期时间将由节点自动设置为该节点最新块的时间戳加上60秒的值。 过期时间间隔可以在节点的配置文件中修改，最大值不能超过24小时。
		tx.Transaction.RawData.Expiration += nonce
		// txHash 一直出错的原因 https://github.com/fbsobreira/gotron-sdk/issues/93
		TronProvider.Client.GRPC.UpdateHash(tx)

		// 对数据签名
		signTx, err := tronrpc.SignTransaction(tx.Transaction, privateKey)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("SignTransaction")
			continue
		}
		// 转换原始数据
		rawData, err := proto.Marshal(signTx)
		if err != nil {
			h.ioc.Logger().Error().Err(err).Msg("Marshal")
			continue
		}
		txHex := hex.EncodeToString(rawData)
		txHash := common.Bytes2Hex(tx.GetTxid())

		for rowIndex, rowID := range info.IDs {
			sendHashMap[rowID] = txHash
			// 只有第一条数据需要发送，其余数据为占位数据
			if rowIndex == 0 {
				sendRows = append(sendRows, &ent.Transfer{
					ChainID:      cfg.ID,
					RelatedType:  transfer.RelatedTypeWithdraw,
					RelatedID:    rowID,
					TxID:         txHash,
					FromAddress:  fromAddress,
					ToAddress:    toAddress,
					AmountStr:    amountStr,
					AmountRaw:    amountRaw,
					Gas:          0,
					GasPrice:     0,
					Nonce:        nonce,
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
					FromAddress:  fromAddress,
					ToAddress:    toAddress,
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
		}
	}

	return sendRows, sendHashMap
}

// 新增数据
func (h *Withdraw) AddWithdraw(sendRows []*ent.Transfer, sendHashMap map[uint64]string) error {

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
