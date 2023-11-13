package eth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/skyisboss/pay-system/ent"

	// "github.com/skyisboss/pay-system/internal/app/kms"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/service/balance"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type DepositParam struct {
	ctx    context.Context
	ioc    *ioc.Container
	abi    abi.ABI
	cfg    *ent.Blockchain
	rpcNum int64
}
type DepositResult struct{}

// eth检测用户充值
// 处理区块数据后写入txn表
// 根据币种配置，处理所有属区块主币和代币交易信息，
// 当配置里 status=0时暂停处理，这里以后需要优化
// 暂停时记录当前区块数，恢复后则从暂停数继续遍历到最新区块，不然会造成暂停期间的入账丢单
func (p *Provider) CheckDeposit() {
	log := p.ioc.Logger()
	ctx := context.Background()
	p.ioc.ApprunService().Lock(ctx, "CheckDeposit")
	defer p.ioc.ApprunService().UnLock(ctx, "CheckDeposit")

	// 获取币种配置信息
	cfgs, err := p.ioc.BlockchainService().GetByChain(ctx, wallet.ETH)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}
	chainIDMap := util.KeyFunc(cfgs, func(e *ent.Blockchain) ChainID {
		return ChainID(e.ID)
	})

	// 获取最新区块信息
	rpcNum, err := p.RpcGetBlockNumber(ctx)
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}

	var wg sync.WaitGroup
	for id, chainCfg := range chainIDMap {
		if chainCfg.Status != 1 {
			log.Warn().Str("chain-type", chainCfg.Types).Str("chain-id", ChainID(id).String()).Msg("系统关闭检测")
			continue
		}

		wg.Add(1)
		var p = &DepositParam{
			ctx:    ctx,
			ioc:    p.ioc,
			abi:    abi.ABI{},
			cfg:    chainCfg,
			rpcNum: rpcNum,
		}
		if chainCfg.TokenAbi == "" && chainCfg.TokenAddress == "" { // 主币 合约地址和abi都为空
			go func() {
				HandleDepositCoin(p)
				wg.Done()
			}()
		} else { // 代币
			// 读取abi
			contractAbi, err := abi.JSON(strings.NewReader(chainCfg.TokenAbi))
			if err != nil {
				log.Info().Err(err).Msg("读取abi")
				continue
			}
			p.abi = contractAbi
			go func() {
				HandleDepositToken(p)
				wg.Done()
			}()
		}
	}
	wg.Wait()

}

// 处理主币
func HandleDepositCoin(p *DepositParam) (*DepositResult, error) {
	// 遍历解析区块信息
	// startIndex := p.startIndex
	// endIndex := p.endIndex
	// startIndex := p.cfg.ScanBlockNum + 1
	// endIndex := p.rpcNum - p.cfg.MinConfirmNum + 1

	// 手续费热钱包，如果块里交易信息to地址是否包含该地址则跳过该记录
	hotWallet := []Address{}

	// 解析区块信息
	parseRpcBlock := func(blockNum int64) ([]Address, map[Address][]*types.Transaction, error) {
		// rpc获取block信息
		rpcBlock, err := p.ioc.EthClient().BlockByNumber(p.ctx, big.NewInt(blockNum))
		if err != nil {
			log.Error().Err(err).Msg("BlockByNumber")
			return nil, nil, err
		}

		// 接收地址列表
		var addressRows []Address
		// map=[toAddress] =>Transaction交易信息
		addressMap := make(map[Address][]*types.Transaction)
		for _, rpcTx := range rpcBlock.Transactions() {
			// 只处理转账数额大于 0 并且非合约的交易
			if rpcTx.Value().Cmp(big.NewInt(0)) <= 0 || rpcTx.To() == nil {
				continue
			}

			from, err := types.Sender(types.LatestSignerForChainID(rpcTx.ChainId()), rpcTx)
			if err != nil {
				return nil, nil, err
			}
			// 如果打币地址在手续费热钱包地址则不处理
			if util.IsStringInSlice(hotWallet, Address(util.AddressBytesToStr(from))) {
				// ioc.Logger().Debug().Msgf("匹配手续费热钱包地址 %s", toAddress)
				continue
			}
			toAddress := util.AddressBytesToStr(*(rpcTx.To()))
			addressMap[Address(toAddress)] = append(addressMap[Address(toAddress)], rpcTx)
			if !util.IsStringInSlice(addressRows, Address(toAddress)) {
				addressRows = append(addressRows, Address(toAddress))
			}
		}
		return addressRows, addressMap, nil
	}

	startIndex, endIndex := makeScanBlockNumber(p.cfg.ScanBlockNum+1, p.rpcNum-p.cfg.MinConfirmNum+1)
	for index := startIndex; index < endIndex; index++ {
		// rpc获取block信息
		toRows, toMap, err := parseRpcBlock(index)
		if err != nil {
			log.Error().Err(err).Msg("HandleBlockData")
			return nil, err
		}

		// 从db中查询这些地址是否是冲币地址
		addressRows, err := p.ioc.AddressService().ListByAddressIn(p.ctx, util.MapSlice(toRows, func(a Address) string {
			return a.String()
		}))
		if err != nil {
			log.Error().Err(err).Msg("ListByAddressIn")
			return nil, err
		}

		dbTxn, err := p.HandleTxnCoin(addressRows, toMap)
		if err != nil {
			log.Error().Err(err).Msg("HandleDepositData")
			return nil, err
		}

		if len(dbTxn) > 0 {
			err = p.HandleDepositBalance(dbTxn)
			if err != nil {
				log.Error().Err(err).Msg("HandleDepositBalance")
				return nil, err
			}
		}

		// 更新检查到的最新区块数到数据库
		err = p.ioc.BlockchainService().UpdateScanBlockByID(p.ctx, p.cfg.ID, index)
		if err != nil {
			log.Error().Err(err).Msg("UpdateScanBlockByID")
			return nil, err
		}
	}

	return nil, nil
}

// 处理代币
func HandleDepositToken(p *DepositParam) (*DepositResult, error) {
	ctx := p.ctx
	cfg := p.cfg
	abi := p.abi
	ioc := p.ioc

	// startIndex := p.cfg.ScanBlockNum + 1
	// endIndex := p.rpcNum - p.cfg.MinConfirmNum + 1
	// 测试模式，只遍历一个区块
	// endIndex := startIndex + 1

	// 手续费热钱包，如果块里交易信息to地址是否包含该地址则跳过该记录
	// hotWallet := []Address{cfg.HotAddress}
	hotWallet := []Address{}

	// 根据区块数获取日志信息
	rpcBlockFilterAddress := func(index int64) ([]Address, map[Address][]types.Log, error) {
		var warpAddresses []common.Address
		event := abi.Events["Transfer"]
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(index),
			ToBlock:   big.NewInt(index),
			Addresses: warpAddresses,
			Topics: [][]common.Hash{
				{event.ID},
			},
		}
		rpcLogs, err := p.ioc.EthClient().FilterLogs(ctx, query)
		if err != nil {
			return nil, nil, err
		}

		// 接收地址列表
		var addressRows []Address
		// map[接收地址] => []交易信息
		addressMap := make(map[Address][]types.Log)

		for _, log := range rpcLogs {
			if log.Removed {
				continue
			}

			toAddress := Address(util.HashToAddrssStringLower(log.Topics[2]))
			// 如果打币地址在手续费热钱包地址则不处理
			if util.InArray(hotWallet, toAddress) {
				ioc.Logger().Debug().Msgf("匹配手续费热钱包地址 %s", toAddress)
				continue
			}
			if !util.IsStringInSlice(addressRows, toAddress) {
				addressRows = append(addressRows, toAddress)
			}
			addressMap[toAddress] = append(addressMap[toAddress], log)
		}

		return addressRows, addressMap, nil
	}
	// 遍历解析区块信息
	startIndex, endIndex := makeScanBlockNumber(p.cfg.ScanBlockNum+1, p.rpcNum-p.cfg.MinConfirmNum+1)
	for index := startIndex; index < endIndex; index++ {
		// rpc获取block信息
		toRows, logsMap, err := rpcBlockFilterAddress(index)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return nil, err
		}

		// 从db中查询这些地址是否是冲币地址
		addressRows, err := p.ioc.AddressService().ListByAddressIn(ctx, util.MapSlice(toRows, func(a Address) string {
			return a.String()
		}))
		if err != nil {
			log.Error().Err(err).Msg("err")
			return nil, err
		}

		dbTxn, err := p.HandleTxnToken(addressRows, logsMap)
		// dbTxn, err := NewDepositToken(cfg, abi, p.ioc, addressRows, logsMap)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return nil, err
		}

		if len(dbTxn) > 0 {
			err = p.HandleDepositBalance(dbTxn)
			if err != nil {
				log.Error().Err(err).Msg("HandleDepositBalance")
				return nil, err
			}
		}

		// 更新检查到的最新区块数到数据库
		err = p.ioc.BlockchainService().UpdateScanBlockByID(ctx, cfg.ID, index)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return nil, err
		}
	}
	return nil, nil
}

// 构建区块遍历
func makeScanBlockNumber(start, end int64) (int64, int64) {
	// return start, end
	return start, start + 1 // 测试模式
}

// 处理主币区块信息
func (d *DepositParam) HandleTxnCoin(rows []*ent.Addres, toMap map[Address][]*types.Transaction) ([]*ent.Txn, error) {
	cfg := d.cfg

	// 待插入数据
	var txnRows []*ent.Txn
	// map[address] => UseTo
	productMap := make(map[Address]int64)

	// 遍历数据库中有交易的地址
	for _, itemRow := range rows {
		// 小于0是系统使用作为热钱包标记
		if itemRow.UseTo < 0 {
			continue
		}

		// 获取 UseTag
		productMap[Address(itemRow.Address)] = itemRow.UseTo

		txes, ok := toMap[Address(itemRow.Address)]
		if !ok {
			return nil, errors.New("err: in blockMap")
		}

		// 获取地址对应的交易列表
		for _, tx := range txes {
			from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				return nil, err
			}
			fromAddress := util.AddressBytesToStr(from)
			toAddress := util.AddressBytesToStr(*(tx.To()))
			// 单位转换
			amountRaw, err := util.BigintToDecimal(tx.Value())
			if err != nil {
				return nil, err
			}
			amountStr, err := util.WeiToEth(tx.Value(), int32(cfg.Decimals))
			if err != nil {
				return nil, err
			}

			txnRows = append(txnRows, &ent.Txn{
				ChainID:      cfg.ID,
				ProductID:    productMap[Address(toAddress)],
				TxID:         tx.Hash().String(),
				FromAddress:  fromAddress,
				ToAddress:    toAddress,
				AmountStr:    amountStr,
				AmountRaw:    amountRaw,
				HandleStatus: txn.HandleStatusInit.Int64(),
			})
		}
	}

	return txnRows, nil
}

// 处理代币区块信息
func (d *DepositParam) HandleTxnToken(rows []*ent.Addres, logsMap map[Address][]types.Log) ([]*ent.Txn, error) {
	cfg := d.cfg
	ioc := d.ioc
	abi := d.abi

	productMap := make(map[Address]int64)
	for _, itemRow := range rows {
		productMap[Address(itemRow.Address)] = itemRow.UseTo
	}

	// 待插入数据
	var dbTxnRows []*ent.Txn
	// 遍历数据库中有交易的地址
	for _, itemRow := range rows {
		// 小于0是系统使用作为热钱包标记, 0是未使用
		if itemRow.UseTo <= 0 {
			continue
		}
		// 获取地址对应的交易列表
		logs, ok := logsMap[Address(itemRow.Address)]
		if !ok {
			ioc.Logger().Error().Msgf("toAddressLogMap no: %s", itemRow.Address)
			return nil, fmt.Errorf("toAddressLogMap no: %s", itemRow.Address)
		}

		// log结构，这里要根据 transactionHash 获取该交易的状态，
		// {
		// 	"address": "0x48cfd205698120a01fec4481cfc5e8a205677791",
		// 	"topics": [
		// 		"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		// 		"0x00000000000000000000000006edba1bf11ff69bfaaec69cf6fc614856989272",
		// 		"0x00000000000000000000000006edba1bf11ff69bfaaec69cf6fc614856989272"
		// 	],
		// 	"data": "0x00000000000000000000000000000000000000000000000000000000054c3380",
		// 	"blockNumber": "0x4526b9",
		// 	"transactionHash": "0x4358c01d5752a79449f6f6f2e58025f35c2d46d67b5e3eb5cad74c546213b612",
		// 	"transactionIndex": "0xb",
		// 	"blockHash": "0x3b1e9c93bc0030a23cc9cf2dd430cd8d58aba27b488007b8173d13b95d9a2124",
		// 	"logIndex": "0x14",
		// 	"removed": false
		// }
		for _, log := range logs {
			// 交易地址
			fromAddress := util.HashToAddrssStringLower(log.Topics[1])
			toAddress := util.HashToAddrssStringLower(log.Topics[2])
			// 提取交易金额
			amount, ok := new(big.Int).SetString(new(big.Int).SetBytes(log.Data).String(), 10)
			if !ok {
				return nil, errors.New("提取交易金额错误")
			}
			// 金额是否大于0
			if amount.Cmp(new(big.Int)) <= 0 {
				return nil, errors.New("error amount <= 0")
			}

			// 判断合约地址是否一致
			abiAddress := strings.ToLower(log.Address.String())
			if abiAddress != cfg.TokenAddress {
				ioc.Logger().Error().
					Str("log.Address", abiAddress).
					Str("TokenAddress", cfg.TokenAddress).
					Msg("合约地址不匹配")
				continue
			}
			// 查询交易状态
			rpcTxReceipt, err := ioc.EthClient().TransactionReceipt(
				context.Background(),
				log.TxHash,
			)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("查询交易回执")
				return nil, err
			}
			// Status： 成功与否，1表示成功，0表示失败
			if rpcTxReceipt.Status != 1 {
				continue
			}

			// 检测收据
			rpcTx, _, err := ioc.EthClient().TransactionByHash(
				context.Background(),
				log.TxHash,
			)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("err")
				return nil, err
			}

			// 检测input
			input, err := abi.Pack("transfer", util.HashToAddrss(log.Topics[2]), amount)
			if err != nil {
				return nil, err
			}
			if hexutil.Encode(input) != hexutil.Encode(rpcTx.Data()) {
				// input 不匹配
				continue
			}

			// 转换金额到eth
			amountStr, err := util.WeiToEth(amount, int32(cfg.Decimals))
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("err")
				return nil, err
			}

			amountRaw, err := util.BigintToDecimal(amount)
			if err != nil {
				return nil, err
			}

			// EnableDeposit := int64(0)
			// if amountRaw.Cmp(cfg.MinDeposit) >= 0 {
			// 	EnableDeposit = 1
			// }
			// EnableCollect := int64(0)
			// if amountRaw.Cmp(cfg.MinCollect) >= 0 {
			// 	EnableCollect = 1
			// }

			dbTxnRows = append(dbTxnRows, &ent.Txn{
				ChainID:      cfg.ID,
				ProductID:    productMap[Address(toAddress)],
				TxID:         log.TxHash.Hex(),
				FromAddress:  fromAddress,
				ToAddress:    toAddress,
				AmountStr:    amountStr,
				AmountRaw:    amountRaw,
				HandleStatus: txn.HandleStatusInit.Int64(),
				// EnableDeposit: EnableDeposit,
				// EnableCollect: EnableCollect,
			})
		}
	}

	return dbTxnRows, nil
}

// 处理用户余额
func (p *DepositParam) HandleDepositBalance(txRows []*ent.Txn) error {
	ctx := p.ctx

	// 开启事务
	tx, err := p.ioc.DBClient().Tx(p.ctx)
	if err != nil {
		p.ioc.Logger().Error().Err(err).Msg("new transactional client")
		return err
	}
	// 事务回滚
	rollback := func(tx *ent.Tx, err error) error {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return err
	}

	res, err := p.ioc.TxnService().CreateManyWithTx(ctx, tx, txRows)
	if err != nil {
		p.ioc.Logger().Error().Err(err).Msg("err")
		return rollback(tx, err)
	}

	// 更新余额
	for _, v := range res {
		_, err = p.ioc.BalanceService().TxUpdateDepositAmount(p.ctx, tx, &balance.UpdateBalance{
			ChainID:   v.ChainID,
			ProductID: uint64(v.ProductID),
			Amount:    v.AmountRaw,
		})
		if err != nil {
			p.ioc.Logger().Error().Err(err).Msg("err")
			return rollback(tx, err)
		}
	}

	// 处理充值通知
	notifyIDs, notifyRows, err := p.CreateDepositNotify(res)
	if err != nil {
		return rollback(tx, err)
	}
	if len(notifyIDs) > 0 {
		ids := util.MapSlice(notifyIDs, func(e PkID) uint64 {
			return uint64(e)
		})
		_, err := p.ioc.TxnService().TxUpdateNotifyStatusByIds(ctx, tx, ids)
		if err != nil {
			log.Error().Err(err).Msg("更新通知状态")
			return rollback(tx, err)
		}
	}
	if len(notifyRows) > 0 {
		_, err := p.ioc.NotifyService().TxCreateMany(ctx, tx, notifyRows)
		if err != nil {
			log.Error().Err(err).Msg("添加通知数据")
			return rollback(tx, err)
		}
	}

	// 提交事物
	return tx.Commit()
}

// 处理充值通知
func (p *DepositParam) CreateDepositNotify(txRows []*ent.Txn) ([]PkID, []*ent.Notify, error) {
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
	productIDMap := util.KeyFunc(productRows, func(c *ent.Product) PkID {
		return PkID(c.ID)
	})

	// 通知id集合
	var notifyIDs []PkID
	// 通知数据
	var notifyRows []*ent.Notify

	for _, txRow := range txRows {
		productRow, ok := productIDMap[PkID(txRow.ProductID)]
		if !ok {
			log.Error().Err(err).Msgf("no productMap: %d", txRow.ProductID)
			notifyIDs = append(notifyIDs, PkID(txRow.ID))
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
		notifyIDs = append(notifyIDs, PkID(txRow.ID))
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
