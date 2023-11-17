package tron

import (
	"context"
	"time"

	"github.com/skyisboss/pay-system/ent"
	ent_transfer "github.com/skyisboss/pay-system/ent/transfer"
	ent_withdraw "github.com/skyisboss/pay-system/ent/withdraw"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

// 检测交易是否上链
// 检查 Transfer 表数据的tx_id 以便确认此交易已经上链
// 根据 related_type 判断是那种交易 1xx - 零钱整理 / 2xx - 用户提币，更新对应数据
func (p *Provider) CheckTxConfirm() {
	log := p.Ioc().Logger()
	ctx := context.Background()
	if ok := p.Ioc().ApprunService().Lock(ctx, "CheckTxConfirm-tron"); !ok {
		return
	}
	defer p.Ioc().ApprunService().UnLock(ctx, "CheckTxConfirm-tron")

	// 获取币种配置信息
	cfgs, err := p.Ioc().BlockchainService().GetByChain(ctx, wallet.TRON)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}
	chainIDs := util.MapSlice(cfgs, func(e *ent.Blockchain) uint64 {
		return e.ID
	})
	chainIDMap := util.KeyFunc(cfgs, func(e *ent.Blockchain) wallet.ChainID {
		return wallet.ChainID(e.ID)
	})

	// 获取待检测数据
	transferRows, err := p.Ioc().TransferService().ListByHandleStatus(ctx, transfer.QuerySend{
		ChainIDs:     chainIDs,
		HandleStatus: transfer.HandleStatusSend,
	})
	if err != nil {
		log.Error().Err(err).Msg("获取数据")
		return
	}
	if len(transferRows) <= 0 {
		log.Info().Err(err).Msg("暂无处理数据")
		return
	}

	util.Println(chainIDMap)
	util.ToJson(transferRows)

	// 获取当前链上最新区块数
	rpcNum, err := p.Ioc().WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider).RpcBlockNumber()
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}

	var relatedIDs []uint64
	relatedTypeMap := make(map[int64][]uint64)
	confirmIDMap := make(map[uint64]*ConfirmUpdates)
	// 根据哈希获取交易信息
	for _, itemRow := range transferRows {
		cfg, ok := chainIDMap[wallet.ChainID(itemRow.ChainID)]
		if !ok {
			log.Error().Msg("无法获取配置")
			continue
		}

		// 占位数据标注为确认
		if itemRow.Nonce < 0 && itemRow.Hex == "" {
			if !util.InArray(relatedIDs, itemRow.RelatedID) {
				relatedIDs = append(relatedIDs, itemRow.RelatedID)
				relatedTypeMap[itemRow.RelatedType] = relatedIDs
			}
			confirmIDMap[itemRow.ID] = &ConfirmUpdates{
				Gas:      0,
				GasPrice: 0,
			}
			continue
		}

		// 根据哈希获取交易信息
		util.Println(rpcNum, cfg)
		rpcTx, err := p.Ioc().TronClient().GRPC.GetTransactionInfoByID(itemRow.TxID)
		if err != nil {
			log.Error().Err(err).Msg("GetTransactionInfoByID")
			continue
		}
		if rpcTx == nil {
			continue
		}
		// 最小确认数
		conrifmNum := rpcNum - rpcTx.BlockNumber + 1
		if conrifmNum < cfg.MinConfirmNum {
			continue
		}

		if !util.InArray(relatedIDs, itemRow.RelatedID) {
			relatedIDs = append(relatedIDs, itemRow.RelatedID)
			relatedTypeMap[itemRow.RelatedType] = relatedIDs
		}
		confirmIDMap[itemRow.ID] = &ConfirmUpdates{
			Gas:      rpcTx.Fee,
			GasPrice: rpcTx.Fee,
		}
	}

	// 开启事务
	tx, err := p.Ioc().DBClient().Tx(ctx)
	if err != nil {
		p.Ioc().Logger().Error().Err(err).Msg("new transactional client")
		return
	}

	// 更新gas
	for i, v := range confirmIDMap {
		id := i
		data := v
		_, err = tx.Transfer.Update().
			Where(ent_transfer.IDEQ(id)).
			Where(ent_transfer.HandleStatusEQ(transfer.HandleStatusSend.Int64())).
			SetGas(data.Gas).
			SetGasPrice(data.GasPrice).
			SetHandleMsg(transfer.HandleStatusConfirm.String()).
			SetHandleStatus(transfer.HandleStatusConfirm.Int64()).
			SetHandleTime(time.Now()).
			Save(ctx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	// 更新状态
	for related, itemIDs := range relatedTypeMap {
		ids := itemIDs
		relatedType := related
		switch relatedType {
		// 零钱整理
		case transfer.RelatedTypeCollect, transfer.RelatedTypeCollectErc20:
			// 零钱整理
			// _, err := tx.Txn.Update().
			// 	Where(ent_txn.IDIn(ids...)).
			// 	SetCollectStatus(txn.CollectStatusDone.Int64()).
			// 	SetCollectMsg(txn.CollectStatusDone.String()).
			// 	SetCollectTime(time.Now()).
			// 	Save(ctx)
			// if err != nil {
			// 	tx.Rollback()
			// 	return
			// }
		// 用户提款
		case transfer.RelatedTypeWithdraw, transfer.RelatedTypeWithdrawErc20:
			_, err := tx.Withdraw.Update().
				Where(ent_withdraw.IDIn(ids...)).
				Where(ent_withdraw.HandleStatusEQ(withdraw.WithdrawStatusSend.Int64())).
				SetHandleStatus(withdraw.WithdrawStatusDone.Int64()).
				SetHandleMsg(withdraw.WithdrawStatusDone.String()).
				SetHandleTime(time.Now()).
				Save(ctx)
			if err != nil {
				tx.Rollback()
				return
			}

			// TODO::通知数据
		}
	}

	tx.Commit()
}

type ConfirmUpdates struct {
	Gas      int64
	GasPrice int64
}
