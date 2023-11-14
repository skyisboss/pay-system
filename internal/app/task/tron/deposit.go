package tron

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/service/balance"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type HandleDeposit struct {
	ctx context.Context
	ioc *ioc.Container
	cfg *ent.Blockchain
}

// 提取的区块交易信息
type BlockTx struct {
	// 交易id
	txid string
	// 交易金额
	amountRaw decimal.Decimal
	amountStr string
	// 发送地址
	fromAddress string
	// 接收地址
	toAddress string
}

// tron 检测用户充值
func (p *Provider) CheckDeposit() {
	log := p.Ioc().Logger()
	ctx := context.Background()
	p.Ioc().ApprunService().Lock(ctx, "CheckDeposit-tron")
	defer p.Ioc().ApprunService().UnLock(ctx, "CheckDeposit-tron")

	// 获取币种配置信息
	cfgs, err := p.Ioc().BlockchainService().GetByChain(ctx, wallet.TRON)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}

	for _, cfg := range cfgs {
		if cfg.Status != 1 {
			log.Warn().
				Str("chain-type", cfg.Types).
				Str("chain-id", wallet.ChainID(cfg.ID).String()).
				Msg("系统关闭检测")
			continue
		}

		handler := &HandleDeposit{
			ctx: ctx,
			ioc: p.Ioc(),
			cfg: cfg,
		}
		if cfg.TokenAbi == "" && cfg.TokenAddress == "" {
			handler.DepositCoin()
		} else {
			handler.DepositToken()
		}
	}
}

func (d *HandleDeposit) DepositToken() error {
	cfg := d.cfg
	ioc := d.ioc
	ctx := d.ctx
	rpcNum, err := ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider).RpcBlockNumber()
	if err != nil {
		log.Error().Err(err).Msg("err")
		return err
	}

	// startIndex := cfg.ScanBlockNum + 1
	// endIndex := rpcNum - cfg.MinConfirmNum + 1
	startIndex, endIndex := makeScanBlockNumber(cfg.ScanBlockNum+1, rpcNum-cfg.MinConfirmNum+1)
	TronProvider := ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)

	// 🌈提取rpc获取的block里交易信息
	for index := startIndex; index < endIndex; index++ {
		// rpc获取block信息
		rpcBlock, err := TronProvider.Client.GRPC.GetBlockByNum(index)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("GetBlockByNum")
			return err
		}

		// 接收地址列表
		var toAddressRows []string
		// 一个地址可能存在多条充值记录
		// {
		// 	address: [BlockTx,BlockTx...]
		// }
		toAddressMap := make(map[string][]*BlockTx)

		// 解析区块里交易记录
		for _, rpcTx := range rpcBlock.Transactions {
			for _, contract := range rpcTx.GetTransaction().GetRawData().GetContract() {
				// Transaction_Contract_TriggerSmartContract 智能合约转账记录
				if contract.Type == core.Transaction_Contract_TriggerSmartContract {
					var tsrc core.TriggerSmartContract
					if err = contract.GetParameter().UnmarshalTo(&tsrc); err != nil { // transform `any` back to Struct
						ioc.Logger().Error().Err(err).Msg("UnmarshalTo")
						return err
					}

					fromAddress := address.HexToAddress(common.BytesToHexString(tsrc.OwnerAddress))
					contractAddress := address.HexToAddress(common.BytesToHexString(tsrc.ContractAddress))
					// 判断是否合约地址
					if contractAddress.String() != cfg.TokenAddress {
						ioc.Logger().Error().Msg("合约地址不匹配")
						continue
					}
					// 解析 input data
					input := tsrc.GetData()
					if len(input) < 68 {
						ioc.Logger().Error().Msg("无法解析 input data")
						continue
					}
					// Transfer函数标识符 的签名
					if common.BytesToHexString(input[:4]) != "0xa9059cbb" {
						ioc.Logger().Error().Msg("Transfer函数标识符不匹配")
						continue
					}
					// 加多1位41 方便截取
					input[15] = 65 // 0x41
					toAddress := address.HexToAddress(common.BytesToHexString(input[15:36]))
					amountBigInt := new(big.Int).SetBytes(common.TrimLeftZeroes(input[36:68]))
					amountRaw := decimal.NewFromInt(amountBigInt.Int64())
					amountStr, err := util.WeiToEth(amountBigInt, int32(cfg.Decimals))
					if err != nil {
						ioc.Logger().Error().Err(err).Msg("WeiToEth")
						continue
					}
					// 判断金额是否负数
					if amountRaw.Cmp(decimal.NewFromInt(0)) <= 0 {
						ioc.Logger().Error().Msg("金额负数")
						continue
					}

					if !util.InArray(toAddressRows, toAddress.String()) {
						toAddressRows = append(toAddressRows, toAddress.String())
					}
					toAddressMap[toAddress.String()] = append(toAddressMap[toAddress.String()], &BlockTx{
						txid:        hex.EncodeToString(rpcTx.GetTxid()),
						amountRaw:   amountRaw,
						amountStr:   amountStr,
						fromAddress: fromAddress.String(),
						toAddress:   toAddress.String(),
					})
				}
			}
		}

		// 新增交易记录和通知数据
		err = d.AddDepositData(toAddressRows, toAddressMap)
		if err != nil {
			return err
		}

		// 更新检查到的最新区块数到数据库
		err = ioc.BlockchainService().UpdateScanBlockByID(ctx, cfg.ID, index)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return err
		}
	}

	return nil
}
func (d *HandleDeposit) DepositCoin() error {
	cfg := d.cfg
	ioc := d.ioc
	ctx := d.ctx
	log := ioc.Logger()
	rpcNum, err := ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider).RpcBlockNumber()
	if err != nil {
		log.Error().Err(err).Msg("err")
		return err
	}

	// startIndex := cfg.ScanBlockNum + 1
	// endIndex := rpcNum - cfg.MinConfirmNum + 1
	startIndex, endIndex := makeScanBlockNumber(cfg.ScanBlockNum+1, rpcNum-cfg.MinConfirmNum+1)
	TronProvider := ioc.WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)

	// 🌈提取rpc获取的block里交易信息
	for index := startIndex; index < endIndex; index++ {
		log.Debug().
			Str("chain", cfg.Chain).
			Str("type", cfg.Types).
			Str("blockNum", util.IntToString(index)).
			Msg("遍历区块")
		// rpc获取block信息
		rpcBlock, err := TronProvider.Client.GRPC.GetBlockByNum(index)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("GetBlockByNum")
			return err
		}

		// 接收地址列表
		var toAddressRows []string
		// 一个地址可能存在多条充值记录
		// {
		// 	address: [BlockTx,BlockTx...]
		// }
		toAddressMap := make(map[string][]*BlockTx)

		// 解析区块里交易记录
		for _, rpcTx := range rpcBlock.Transactions {
			for _, contract := range rpcTx.GetTransaction().GetRawData().GetContract() {
				// Transaction_Contract_TransferContract 主币转账记录
				if contract.Type == core.Transaction_Contract_TransferContract {
					var tsrc core.TransferContract
					if err = contract.GetParameter().UnmarshalTo(&tsrc); err != nil { // transform `any` back to Struct
						ioc.Logger().Error().Err(err).Msg("UnmarshalTo")
						return err
					}

					fromAddress := address.HexToAddress(common.BytesToHexString(tsrc.OwnerAddress))
					toAddress := address.HexToAddress(common.BytesToHexString(tsrc.ToAddress))

					// 如果打币地址在手续费热钱包地址则不处理
					if cfg.HotAddress == fromAddress.String() {
						continue
					}

					amountBigInt := new(big.Int).SetInt64(tsrc.Amount)
					amountRaw := decimal.NewFromInt(amountBigInt.Int64())
					amountStr, err := util.WeiToEth(amountBigInt, int32(cfg.Decimals))
					if err != nil {
						ioc.Logger().Error().Err(err).Msg("WeiToEth")
						continue
					}
					// 判断金额是否负数
					if amountRaw.Cmp(decimal.NewFromInt(0)) <= 0 {
						ioc.Logger().Error().Msg("金额负数")
						continue
					}
					if !util.InArray(toAddressRows, toAddress.String()) {
						toAddressRows = append(toAddressRows, toAddress.String())
					}
					toAddressMap[toAddress.String()] = append(toAddressMap[toAddress.String()], &BlockTx{
						txid:        hex.EncodeToString(rpcTx.GetTxid()),
						amountRaw:   amountRaw,
						amountStr:   amountStr,
						fromAddress: fromAddress.String(),
						toAddress:   toAddress.String(),
					})
				}
			}
		}

		// 新增交易记录和通知数据
		err = d.AddDepositData(toAddressRows, toAddressMap)
		if err != nil {
			return err
		}

		// 更新检查到的最新区块数到数据库
		err = ioc.BlockchainService().UpdateScanBlockByID(ctx, cfg.ID, index)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return err
		}
	}

	return nil
}

// 添加充值数据
func (d *HandleDeposit) AddDepositData(toAddressRows []string, toAddressMap map[string][]*BlockTx) error {
	cfg := d.cfg
	ioc := d.ioc
	ctx := d.ctx

	// 1、从db中查询这些地址是否是冲币地址
	toRows, err := ioc.AddressService().ListByAddressIn(ctx, toAddressRows)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("ListByAddressIn")
		return err
	}

	// 待插入数据
	var dbTxnRows []*ent.Txn
	addressProductMap := make(map[string]int64)
	for _, itemRow := range toRows {
		// UseTag小于0是系统使用作为热钱包标记
		if itemRow.UseTo < 0 {
			continue
		}
		// 提取产品id
		addressProductMap[itemRow.Address] = itemRow.UseTo
		txes, ok := toAddressMap[itemRow.Address]
		if !ok {
			ioc.Logger().Error().Msgf("err: %s not in blockMap", itemRow.Address)
			continue
		}
		for _, tx := range txes {
			vtx := tx
			dbTxnRows = append(dbTxnRows, &ent.Txn{
				TxID:         vtx.txid,
				ChainID:      cfg.ID,
				ProductID:    addressProductMap[vtx.toAddress],
				FromAddress:  vtx.fromAddress,
				ToAddress:    vtx.toAddress,
				AmountRaw:    vtx.amountRaw,
				AmountStr:    vtx.amountStr,
				HandleStatus: txn.HandleStatusInit.Int64(),
				HandleMsg:    txn.HandleStatusInit.String(),
				HandleTime:   time.Now(),
			})
		}
	}

	// 🌈插入交易数据
	// 开启事务
	tx, err := ioc.DBClient().Tx(ctx)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("new transactional client")
		return err
	}
	// 事务回滚
	rollback := func(tx *ent.Tx, err error) error {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return err
	}
	if len(dbTxnRows) > 0 {
		res, err := ioc.TxnService().CreateManyWithTx(ctx, tx, dbTxnRows)
		if err != nil {
			ioc.Logger().Error().Err(err).Msg("err")
			return rollback(tx, err)
		}

		// 更新余额
		for _, v := range res {
			_, err = ioc.BalanceService().TxUpdateDepositAmount(ctx, tx, &balance.UpdateBalance{
				ChainID:   v.ChainID,
				ProductID: uint64(v.ProductID),
				Amount:    v.AmountRaw,
			})
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("err")
				return rollback(tx, err)
			}
		}

		// 处理充值通知
		notifyIDs, notifyRows, err := d.AddDepositNotify(res)
		if err != nil {
			return rollback(tx, err)
		}
		if len(notifyIDs) > 0 {
			_, err := ioc.TxnService().TxUpdateNotifyStatusByIds(ctx, tx, notifyIDs)
			if err != nil {
				log.Error().Err(err).Msg("更新通知状态")
				return rollback(tx, err)
			}
		}
		if len(notifyRows) > 0 {
			_, err := ioc.NotifyService().TxCreateMany(ctx, tx, notifyRows)
			if err != nil {
				log.Error().Err(err).Msg("添加通知数据")
				return rollback(tx, err)
			}
		}
	}

	return tx.Commit()
}

// 处理充值通知
func (p *HandleDeposit) AddDepositNotify(txRows []*ent.Txn) ([]uint64, []*ent.Notify, error) {
	ctx := p.ctx
	cfg := p.cfg
	ioc := p.ioc
	productRows, err := ioc.ProductService().ListInIDs(ctx, util.MapSlice(txRows, func(e *ent.Txn) uint64 {
		return uint64(e.ProductID)
	}))
	if err != nil {
		log.Error().Err(err).Msg("err")
		return nil, nil, err
	}
	productIDMap := util.KeyFunc(productRows, func(c *ent.Product) uint64 {
		return c.ID
	})

	// 通知id集合
	var notifyIDs []uint64
	// 通知数据
	var notifyRows []*ent.Notify

	for _, txRow := range txRows {
		productRow, ok := productIDMap[uint64(txRow.ProductID)]
		if !ok {
			log.Error().Err(err).Msgf("no productMap: %d", txRow.ProductID)
			notifyIDs = append(notifyIDs, uint64(txRow.ID))
			continue
		}

		// 构建通知post提交数据
		now := time.Now()
		amountRaw, _ := txRow.AmountRaw.Float64()
		bodyObj := gin.H{
			"appid":       productRow.AppID,
			"address":     txRow.ToAddress,
			"chain":       cfg.Chain,
			"symbol":      fmt.Sprintf("%s-%s", cfg.Types, cfg.Symbol),
			"decimals":    cfg.Decimals,
			"amount_str":  txRow.AmountStr,
			"amount_raw":  amountRaw,
			"tx_hash":     txRow.TxID,
			"notify_type": notify.NotifyItemDeposit.Int(),
			"notify_desc": notify.NotifyItemDeposit.String(),
			"handle_time": now.Unix(),
		}
		bodyObj["sign"] = util.WechatSign(productRow.AppSecret, bodyObj)
		body, err := json.Marshal(bodyObj)
		if err != nil {
			log.Error().Err(err).Msg("err")
			continue
		}
		notifyIDs = append(notifyIDs, uint64(txRow.ID))
		notifyRows = append(notifyRows, &ent.Notify{
			Nonce:        util.GetUUID(),
			ChainID:      cfg.ID,
			ProductID:    uint64(txRow.ProductID),
			ItemType:     notify.NotifyItemDeposit.Int(),
			ItemFrom:     uint64(txRow.ID),
			NotifyType:   notify.NotifyTypeReceived,
			SendURL:      productRow.WebHook,
			SendBody:     string(body),
			HandleStatus: notify.HandleStatusInit.Int(),
			HandleMsg:    notify.HandleStatusInit.String(),
			HandleTime:   now,
		})
	}

	return notifyIDs, notifyRows, nil
}

// 构建区块遍历
func makeScanBlockNumber(start, end int64) (int64, int64) {
	// return start, end
	return start, start + 1 // 测试模式
}
