package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/util"
)

// Collect 待整理信息
type Collect struct {
	// 待整理数据id
	TxnIDs []uint64
	// 发送金额 wei单位
	SendAmount *big.Int
}

// CollectAmountLimit 最小零钱整理金额限制
const CollectAmountLimit float64 = 0.1

// CheckCollect 零钱整理-余额归集 从txn表获取待整理数据，处理后写入transfer表
// 为了资金安全，需要定时将用户充值的金额转移到指定的冷钱包
func (p *Provider) CheckCollect() {
	log := p.ioc.Logger()
	ctx := context.Background()
	blockService := p.ioc.BlockchainService()

	// 获取币种配置
	chainCfg, gas, err := blockService.GetByChainAndType(ctx, "eth", "eth")
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}

	// 获取待整理数据
	addressRows, addressMap, err := getCollectRows(ctx, p.ioc, chainCfg)
	if err != nil {
		return
	}

	// 待整理数据里提取对应私钥
	pkMap, err := getPrivateKey(ctx, p.ioc, addressRows)
	if err != nil {
		return
	}

	gasLimit := int64(21000)
	gasPrice := gas.ColdGasPrice
	feeValue := big.NewInt(gasLimit * gasPrice)
	tipPrice := big.NewInt(gas.TipFeePrice)
	coldWallet := chainCfg.ColdAddress
	NetworkID, _ := p.ioc.EthClient().NetworkID(p.ioc.Context())

	// 处理数据，构建tx交易信息
	for address, collectInfo := range addressMap {
		// 获取私钥
		prv, ok := pkMap[address]
		if !ok {
			continue
		}

		// 获取nonce值
		nonceValue, err := getAddressNonce(ctx, p.ioc, address)
		if nil != err {
			return
		}

		// 发送金额
		sendAmount := new(big.Int).Sub(collectInfo.SendAmount, feeValue)
		// 如果不够支付手续费则跳过这条
		if sendAmount.Cmp(new(big.Int)) <= 0 {
			continue
		}

		// 创建交易
		signedTx, err := NewSignTransaction(&NewTransaction{
			nonce:     nonceValue,
			to:        coldWallet,
			amount:    sendAmount,
			gasLimit:  gasLimit,
			gasFeeCap: gasPrice,
			gasTipCap: tipPrice.Int64(),
			chainID:   NetworkID.Int64(),
			prv:       prv,
		})
		if err != nil {
			log.Error().Err(err).Msg("NewSignTransaction")
			return
		}
		rawTxBytes, err := signedTx.MarshalBinary()
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
		txHex := hex.EncodeToString(rawTxBytes)
		txHash := strings.ToLower(signedTx.Hash().Hex())
		amountRaw, _ := util.BigintToDecimal(sendAmount)
		amountStr, _ := util.WeiToEth(sendAmount, int32(chainCfg.Decimals))

		// 创建存入数据
		var sendRows []*ent.Transfer
		for rowIndex, rowID := range collectInfo.TxnIDs {
			if rowIndex == 0 {
				// 只有第一条数据需要发送，其余数据为占位数据
				sendRows = append(sendRows, &ent.Transfer{
					ChainID: chainCfg.ID,
					// 发送类型
					RelatedType:  transfer.RelatedTypeCollect,
					RelatedID:    rowID,
					TxID:         txHash,
					FromAddress:  address.String(),
					ToAddress:    coldWallet,
					AmountStr:    amountStr,
					AmountRaw:    amountRaw,
					Gas:          gasLimit,
					GasPrice:     gasPrice,
					Nonce:        nonceValue,
					Hex:          txHex,
					HandleStatus: transfer.HandleStatusInit.Int64(),
					HandleMsg:    transfer.HandleStatusInit.String(),
					HandleTime:   time.Now(),
				})
			} else {
				// 占位数据
				sendRows = append(sendRows, &ent.Transfer{
					ChainID:      chainCfg.ID,
					RelatedType:  transfer.RelatedTypeCollect,
					RelatedID:    rowID,
					TxID:         txHash,
					FromAddress:  address.String(),
					ToAddress:    coldWallet,
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

		// 插入待发送数据
		_, err = p.ioc.TransferService().CreateMany(ctx, sendRows)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
		// // 更改tx整理状态
		// _, err = p.ioc.TxnService().UpdateCollectStatusByIds(ctx, collectInfo.TxnIDs, txn.CollectStatusHex)
		// if err != nil {
		// 	log.Error().Err(err).Msg("err")
		// 	return
		// }
	}
}

// CheckCollectErc20 erc转账时需要支付两笔费用 1、eth手续费 2、erc20代币金额
// 当用户向这个地址(0xA)充入某个erc20代币后，这个地址(0xA)只有erc20代币却没有eth余额。
// 如果想把地址(0xA)某个erc20代币整理到冷钱包，则需要先给0xA冲入一定数量的eth作为0xA转移erc20代币时需要的eth手续费。
//
// 实现思路：首先对待整理地址打入手续费，然后再将待整理地址erc20代币转移出来。
// 即：每个erc20整理操作会附带两次转账（1、对整理地址打入手续费，2、整理地址转出erc20到冷钱包 ），
// 因此erc20整理应该设定一个最小整理额度，否则会浪费很多手续费
func (p *Provider) CheckCollectErc20() {
	log := p.ioc.Logger()
	ctx := context.Background()
	blockService := p.ioc.BlockchainService()

	// 获取币种配置
	chainCfg, gas, err := blockService.GetByChainAndType(ctx, "eth", "erc20")
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}

	// 获取待整理数据
	addressRows, addressMap, err := getCollectRows(ctx, p.ioc, chainCfg)
	if err != nil {
		return
	}

	// 待整理数据里提取对应私钥
	pkMap, err := getPrivateKey(ctx, p.ioc, addressRows)
	if err != nil {
		return
	}

	gasLimit := int64(2100000)
	gasPrice := gas.ColdGasPrice
	// feeValue := big.NewInt(gasLimit * gasPrice)
	tipPrice := big.NewInt(gas.TipFeePrice)
	coldWallet := chainCfg.ColdAddress
	NetworkID, err := p.ioc.EthClient().NetworkID(p.ioc.Context())
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(chainCfg.TokenAbi))
	if err != nil {
		log.Error().Err(err).Msg("读取abi")
		return
	}

	// 处理数据，构建tx交易信息
	for address, collectInfo := range addressMap {
		// 发送金额
		sendAmount := collectInfo.SendAmount

		// 获取私钥
		prv, ok := pkMap[address]
		if !ok {
			continue
		}

		// 获取nonce值
		nonceValue, err := getAddressNonce(ctx, p.ioc, address)
		if nil != err {
			return
		}

		// 生成交易
		input, err := contractAbi.Pack(
			"transfer",
			common.HexToAddress(coldWallet),
			sendAmount,
		)
		if err != nil {
			log.Error().Err(err).Msg("contractAbi.Pack")
			return
		}

		// 创建代币转账交易
		signedTx, err := NewSignTransaction(&NewTransaction{
			nonce:     nonceValue,
			to:        chainCfg.TokenAddress, // 发送到合约地址，token转入地址写在input数据里
			amount:    big.NewInt(0),         // token转账时金额为0，因为代币写在input数据里
			gasLimit:  gasLimit,
			gasFeeCap: gasPrice,
			gasTipCap: tipPrice.Int64(),
			chainID:   NetworkID.Int64(),
			prv:       prv,
			data:      input,
		})
		if err != nil {
			log.Error().Err(err).Msg("NewSignTransaction")
			return
		}
		rawTxBytes, err := signedTx.MarshalBinary()
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
		txHex := hex.EncodeToString(rawTxBytes)
		txHash := strings.ToLower(signedTx.Hash().Hex())
		amountRaw, _ := util.BigintToDecimal(sendAmount)
		amountStr, _ := util.WeiToEth(sendAmount, int32(chainCfg.Decimals))

		// 创建存入数据
		var sendRows []*ent.Transfer
		for rowIndex, rowID := range collectInfo.TxnIDs {
			if rowIndex == 0 {
				// 只有第一条数据需要发送，其余数据为占位数据
				sendRows = append(sendRows, &ent.Transfer{
					ChainID: chainCfg.ID,
					// 发送类型
					RelatedType:  transfer.RelatedTypeCollectErc20,
					RelatedID:    rowID,
					TxID:         txHash,
					FromAddress:  address.String(),
					ToAddress:    coldWallet,
					AmountStr:    amountStr,
					AmountRaw:    amountRaw,
					Gas:          gasLimit,
					GasPrice:     gasPrice,
					Nonce:        nonceValue,
					Hex:          txHex,
					HandleStatus: transfer.HandleStatusInit.Int64(),
					HandleMsg:    transfer.HandleStatusInit.String(),
					HandleTime:   time.Now(),
				})
			} else {
				// 占位数据
				sendRows = append(sendRows, &ent.Transfer{
					ChainID:      chainCfg.ID,
					RelatedType:  transfer.RelatedTypeCollectErc20,
					RelatedID:    rowID,
					TxID:         txHash,
					FromAddress:  address.String(),
					ToAddress:    coldWallet,
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

		// 插入待发送数据
		_, err = p.ioc.TransferService().CreateMany(ctx, sendRows)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
		// 更改tx整理状态
		_, err = p.ioc.TxnService().UpdateCollectStatusByIds(ctx, collectInfo.TxnIDs, txn.CollectStatusHex)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
	}
}

// 获取待整理数据
func getCollectRows(ctx context.Context, ioc *ioc.Container, chainCfg *ent.Blockchain) ([]Address, map[Address]*Collect, error) {
	// 最小零钱整理金额限制
	var minCollectAmount = chainCfg.MinWithdraw.Mul(decimal.NewFromFloat(
		math.Pow10(int(chainCfg.Decimals)),
	))

	// 零钱金额大于某个值才整理
	// limit := decimal.NewFromFloat(CollectAmountLimit)
	txRows, err := ioc.TxnService().ListByCollectChain(ctx, chainCfg.ID, minCollectAmount)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("err")
		return nil, nil, err
	}
	if len(txRows) <= 0 {
		// 没有要处理的信息
		err = errors.New("没有要处理的信息")
		ioc.Logger().Info().Msg(err.Error())
		return nil, nil, err
	}

	// addresses 需要整理的地址列表
	var toAddressRow []Address
	// addressMap map[ToAddress] => CollectInfo
	toAddressMap := make(map[Address]*Collect)

	// 将待整理数据按地址做归并处理
	// id	交易id	冲币地址	冲币金额	零钱整理状态
	// 1	0x1		0xa			0.11	0
	// 2	0x2		0xb			0.12	0
	// 3	0x3		0xa			0.08	0
	// 4	0x4		0xc			0.17	0
	// 这里我们可以看到有这样一种情况，id 是1和3的交易记录都是同一个冲币地址的交易记录，
	// 为了节省手续费，我们会把这两个记录一起整理 合并相同冲币地址的记录，记录为map[冲币地址]SUM(冲币金额)

	for _, itemRow := range txRows {
		collect := toAddressMap[Address(itemRow.ToAddress)]
		if collect == nil {
			collect = &Collect{
				TxnIDs:     []uint64{},
				SendAmount: new(big.Int),
			}
			toAddressMap[Address(itemRow.ToAddress)] = collect
		}
		collect.TxnIDs = append(collect.TxnIDs, itemRow.ID)
		// 字符串金额wei按10进制转换为big.Int wei
		amountWei := itemRow.AmountRaw.BigInt()
		collect.SendAmount.Add(collect.SendAmount, amountWei)

		if !util.InArray(toAddressRow, Address(itemRow.ToAddress)) {
			toAddressRow = append(toAddressRow, Address(itemRow.ToAddress))
		}
	}

	return toAddressRow, toAddressMap, nil
}

// 获取私钥
func getPrivateKey(ctx context.Context, ioc *ioc.Container, rows []Address) (map[Address]*ecdsa.PrivateKey, error) {
	privateKeyMap := make(map[Address]*ecdsa.PrivateKey)
	addressRows, err := ioc.AddressService().ListByAddressIn(ctx, util.MapSlice(rows, func(a Address) string {
		return a.String()
	}))
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("err")
		return nil, err
	}

	for _, itemRow := range addressRows {
		// 解密私钥
		saltKey := util.StrReverse(itemRow.Address+itemRow.UUID) + ioc.Config().Providers.SaltKey
		key, err := util.AesDecrypt(itemRow.Password, saltKey)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("err")
			return nil, err
		}
		if len(key) == 0 {
			err := errors.New("key len = 0")
			ioc.Logger().Error().Err(err).Msg(err.Error())
			return nil, err
		}
		key = strings.TrimPrefix(key, "0x")
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("err")
			return nil, err
		}
		privateKeyMap[Address(itemRow.Address)] = privateKey
	}
	return privateKeyMap, nil
}

// 获取nonce值 1、通过rpc获取 2、通过数据库获取 ,返回两者最大值
func getAddressNonce(ctx context.Context, ioc *ioc.Container, address Address) (int64, error) {
	rpcNonce, err := ioc.EthClient().NonceAt(ctx, common.HexToAddress(address.String()), nil)
	if nil != err {
		ioc.Logger().Error().Err(err).Msg("NonceAt")
		return 0, err
	}
	var dbNonce uint64
	dbNum, err := ioc.TransferService().GetNonceByFromAddress(ctx, address.String())
	if nil != err {
		dbNonce = 1
	} else {
		dbNonce = uint64(dbNum + 1)
	}
	if dbNonce > rpcNonce {
		rpcNonce = dbNonce
	}
	return int64(rpcNonce), nil
}

// RpcBalanceAt 获取主币余额
func (p *Provider) RpcBalanceAt(ctx context.Context, address string) (*big.Int, error) {
	balance, err := p.ioc.EthClient().BalanceAt(ctx, common.HexToAddress(address), nil)
	if nil != err {
		return nil, err
	}
	return balance, nil
}

// RpcTokenBalanceAt 获取token余额
func (p *Provider) RpcTokenBalanceAt(ctx context.Context, tokenAbi string, tokenAddress, address string) (*big.Int, error) {
	tokenAddressHash := common.HexToAddress(tokenAddress)
	// 生成交易
	contractAbi, err := abi.JSON(strings.NewReader(tokenAbi))
	if err != nil {
		return nil, err
	}
	input, err := contractAbi.Pack(
		"balanceOf",
		common.HexToAddress(address),
	)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		From:  common.HexToAddress(address),
		To:    &tokenAddressHash,
		Value: nil,
		Data:  input,
	}
	out, err := p.ioc.EthClient().CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	res, err := contractAbi.Unpack("balanceOf", out)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("error call res")
	}

	out0, ok := res[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("error call res")
	}
	return out0, nil
}

type NewTransaction struct {
	nonce     int64
	to        string
	amount    *big.Int
	gasLimit  int64
	gasFeeCap int64
	gasTipCap int64
	chainID   int64
	data      []byte
	prv       *ecdsa.PrivateKey
}

func NewSignTransaction(ntx *NewTransaction) (*types.Transaction, error) {
	toAddress := common.HexToAddress(ntx.to)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(ntx.chainID),
		Nonce:     uint64(ntx.nonce),
		GasTipCap: big.NewInt(ntx.gasTipCap),
		GasFeeCap: big.NewInt(ntx.gasFeeCap),
		Gas:       uint64(ntx.gasLimit),
		To:        &toAddress,
		Value:     ntx.amount,
		Data:      ntx.data,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(ntx.chainID)), ntx.prv)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
