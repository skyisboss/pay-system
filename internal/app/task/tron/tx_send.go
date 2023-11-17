package tron

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gin-gonic/gin"
	"github.com/skyisboss/pay-system/ent"
	ent_transfer "github.com/skyisboss/pay-system/ent/transfer"
	ent_withdraw "github.com/skyisboss/pay-system/ent/withdraw"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
	"google.golang.org/protobuf/proto"
)

// 发送交易【资产操作】
func (p *Provider) CheckTxSend() {
	log := p.Ioc().Logger()
	ctx := context.Background()

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

	// 获取待发送数据
	sendRows, err := p.Ioc().TransferService().ListByHandleStatus(ctx, transfer.QuerySend{
		ChainIDs:     chainIDs, // 相同网络币种统一处理
		HandleStatus: transfer.HandleStatusInit,
	})
	if err != nil {
		log.Error().Err(err).Msg("获取待发送数据")
		return
	}
	if len(sendRows) <= 0 {
		log.Info().Err(err).Msg("没有要处理的数据")
		return
	}

	collectRows := []*ent.Transfer{}
	withdrawRows := []*ent.Transfer{}
	for _, sendRow := range sendRows {
		switch sendRow.RelatedType {
		case transfer.RelatedTypeCollect, transfer.RelatedTypeCollectErc20:
			collectRows = append(collectRows, sendRow)
		case transfer.RelatedTypeWithdraw, transfer.RelatedTypeWithdrawErc20:
			withdrawRows = append(withdrawRows, sendRow)
		}
	}
	handler := &SendParam{
		ctx:    ctx,
		ioc:    p.Ioc(),
		cfg:    &ent.Blockchain{},
		cfgMap: chainIDMap,
	}

	handler.HandleWithdraw(withdrawRows)
	handler.HandleCollect(collectRows)
}

type SendParam struct {
	ctx    context.Context
	ioc    *ioc.Container
	cfg    *ent.Blockchain
	cfgMap map[wallet.ChainID]*ent.Blockchain
}

func (s *SendParam) HandleCollect(sendRows []*ent.Transfer) error {
	s.ioc.Logger().Warn().Msg("TODO::HandleCollect")
	return nil
}
func (s *SendParam) HandleWithdraw(sendRows []*ent.Transfer) error {
	ioc := s.ioc
	ctx := s.ctx
	log := ioc.Logger()

	// RelatedID
	relatedIDs := util.MapSlice(sendRows, func(e *ent.Transfer) uint64 {
		return e.RelatedID
	})
	txRows, err := ioc.WithdrawService().ListByIDsIn(ctx, relatedIDs)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("err")
		return err
	}
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("err")
		return err
	}
	txIDMap := util.KeyFunc(txRows, func(e *ent.Withdraw) uint64 {
		return e.ID
	})
	// 获取产品
	productIDs := util.MapSlice(txRows, func(e *ent.Withdraw) uint64 {
		return uint64(e.ProductID)
	})
	productRows, err := ioc.ProductService().ListInIDs(ctx, productIDs)
	if err != nil {
		ioc.Logger().Error().Err(err).Msg("err")
		return err
	}
	productIDMap := util.KeyFunc(productRows, func(e *ent.Product) uint64 {
		return e.ID
	})
	var withdrawIDs []uint64
	var transferIDs []uint64
	var notifyRows []*ent.Notify
	onSendOk := func(sendRow *ent.Transfer) error {
		// 将发送成功和占位数据计入数组
		if !util.IsIntInSlice(transferIDs, sendRow.ID) {
			transferIDs = append(transferIDs, sendRow.ID)
		}
		if !util.IsIntInSlice(withdrawIDs, uint64(sendRow.RelatedID)) {
			withdrawIDs = append(withdrawIDs, uint64(sendRow.RelatedID))
		}

		txnRow, ok := txIDMap[sendRow.RelatedID]
		if !ok {
			ioc.Logger().Error().Msgf("txnIDMap no: %d", sendRow.RelatedID)
			return nil
		}

		productRow, ok := productIDMap[uint64(txnRow.ProductID)]
		if !ok {
			ioc.Logger().Error().Msgf("productMap no: %d", txnRow.ProductID)
			return nil
		}

		// 创建通知信息
		cfg := s.cfgMap[wallet.ChainID(sendRow.ChainID)]
		nonce := util.GetUUID()
		// 构建通知post提交数据
		symbol := fmt.Sprintf("%s-%s", cfg.Types, cfg.Symbol)
		if cfg.Types == cfg.Symbol {
			symbol = cfg.Symbol
		}
		reqObj := gin.H{
			"appid":       productRow.AppID,
			"address":     txnRow.ToAddress,
			"chain":       cfg.Chain,
			"symbol":      symbol,
			"decimals":    cfg.Decimals,
			"amount_str":  txnRow.AmountStr,
			"amount_raw":  txnRow.AmountRaw.BigInt(),
			"tx_hash":     sendRow.TxID,
			"serial_id":   nonce,
			"notify_type": notify.NotifyItemWithdraw,
			"notify_desc": "send withdraw",
		}
		reqObj["sign"] = util.WechatSign(productRow.AppSecret, reqObj)
		req, err := json.Marshal(reqObj)
		if err != nil {
			log.Error().Err(err).Msg("Marshal")
			return err
		}
		notifyRows = append(notifyRows, &ent.Notify{
			Nonce:        nonce,
			ChainID:      txnRow.ChainID,
			ProductID:    uint64(txnRow.ProductID),
			ItemType:     sendRow.RelatedType,
			ItemFrom:     uint64(txnRow.ID),
			NotifyType:   "send withdraw",
			SendURL:      productRow.WebHook,
			SendBody:     string(req),
			HandleStatus: notify.HandleStatusInit.Int(),
			HandleMsg:    notify.HandleStatusInit.String(),
		})

		return nil
	}

	for _, sendRow := range sendRows {
		// 发送数据中需要排除占位数据
		if sendRow.Hex != "" {
			// 还原16进制数据
			b, _ := hex.DecodeString(sendRow.Hex)
			var rawData core.Transaction
			err = proto.Unmarshal(b, &rawData)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("err")
				continue
			}
			// 广播交易 执行资产转账操作
			result, err := ioc.TronClient().GRPC.Broadcast(&rawData)
			if err != nil {
				ioc.Logger().Error().Err(err).Msg("err")
				continue
			}
			if result.Code != 0 {
				err = errors.New(string(result.GetMessage()))
				ioc.Logger().Error().Err(err).Msg("err")
				continue
			}
		}
		err = onSendOk(sendRow)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return err
		}
	}

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
	if len(withdrawIDs) > 0 {
		_, err := tx.Withdraw.Update().
			Where(ent_withdraw.IDIn(withdrawIDs...)).
			SetHandleStatus(withdraw.WithdrawStatusSend.Int64()).
			SetHandleMsg(withdraw.WithdrawStatusSend.String()).
			SetHandleTime(time.Now()).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}
	if len(transferIDs) > 0 {
		_, err := tx.Transfer.Update().
			Where(ent_transfer.IDIn(transferIDs...)).
			SetHandleStatus(transfer.HandleStatusSend.Int64()).
			SetHandleMsg(transfer.HandleStatusSend.String()).
			SetHandleTime(time.Now()).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}
	if len(notifyRows) > 0 {
		rows := []*ent.NotifyCreate{}
		for _, v := range notifyRows {
			row := tx.Notify.Create().
				SetChainID(v.ChainID).
				SetItemType(v.ItemType).
				SetItemFrom(v.ItemFrom).
				SetNonce(v.Nonce).
				SetProductID(v.ProductID).
				SetNotifyType(v.NotifyType).
				SetSendBody(v.SendBody).
				SetSendURL(v.SendURL).
				SetSendRetry(v.SendRetry).
				SetHandleStatus(v.HandleStatus).
				SetCreatedAt(time.Now())
			rows = append(rows, row)
		}
		_, err := tx.Notify.CreateBulk(rows...).Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	// 提交事物
	return tx.Commit()
}
