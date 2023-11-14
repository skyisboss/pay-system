package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/rpc/webhook"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Service struct {
	ioc     *ioc.Container
	handler *Handler
	logger  *zerolog.Logger
}

func New(
	ioc *ioc.Container,
	handler *Handler,
	logger *zerolog.Logger,
) *Service {
	log := logger.With().Str("channel", "address_task").Logger()

	return &Service{
		ioc:     ioc,
		handler: handler,
		logger:  &log,
	}
}
func (s *Service) EthProvider() *eth.Provider {
	p := s.handler.providers[wallet.ETH].(*eth.Provider)
	return p.SetInit(s.ioc)
}

func (s *Service) TronProvider() *tron.Provider {
	return s.handler.providers[wallet.TRON].(*tron.Provider)
}

func (s *Service) GetProvider(blockchain wallet.Blockchain) Provider {
	return s.handler.providers[blockchain]
}

// 检测剩余可用地址 获取币种配置信息 过滤出主币
func (s *Service) CheckFreeAddress() {
	ioc := s.ioc
	log := ioc.Logger()
	ctx := context.Background()
	if ok := ioc.ApprunService().Lock(ctx, "CheckFreeAddress"); !ok {
		return
	}
	defer ioc.ApprunService().UnLock(ctx, "CheckFreeAddress")

	// 获取币种配置信息
	cfgs, err := ioc.BlockchainService().GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("获取币种配置")
		return
	}

	for _, item := range cfgs {
		cfg := item
		// 过滤代币，只需要主币
		if cfg.Chain != cfg.Types || cfg.TokenAbi != "" || cfg.TokenAddress != "" {
			continue
		}
		// 统计address表里剩余可用数量
		count, err := ioc.AddressService().GetCountByChainID(ctx, cfg.ID)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}

		freeNum := cfg.MinFreeNum - count
		if freeNum <= 0 {
			continue
		}

		var rows []*ent.Addres
		provider := ioc.WalletService().GetProvider(wallet.Blockchain(cfg.Chain))
		for i := int64(0); i < freeNum; i++ {
			wt := provider.CreateWallet()
			if wt == nil {
				log.Error().Msg("CreateWallet error")
				continue
			}
			password, err := wt.EncodePrivateKey(ioc.Config().Providers.SaltKey)
			if err != nil {
				log.Error().Err(err).Msg("EncodePrivateKey")
				continue
			}
			rows = append(rows, &ent.Addres{
				ChainID:  cfg.ID,
				Address:  wt.Address,
				Password: password,
				UUID:     wt.UUID,
			})
		}
		if len(rows) > 0 {
			_, err = ioc.AddressService().CreateMany(ctx, rows)
			if err != nil {
				log.Error().Err(err).Msg("err")
				continue
			}
		}
	}
}

// 检测新增回调通知数据
// txn表：检测充值
// transfer表：检测交易发送
// withdraw表：检测提款信息
func (s *Service) CheckAddNotify() {
	log := s.logger
	ctx := context.Background()
	txnService := s.ioc.TxnService()
	notifyService := s.ioc.NotifyService()
	productService := s.ioc.ProductService()

	// 获取交易数据
	txRows, err := txnService.ListByHandleStatus(ctx, txn.HandleStatusInit.Int64())
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}
	if len(txRows) <= 0 {
		log.Info().Msg("没有要处理的数据")
		return
	}

	// 获取产品数据
	_ids := util.MapSlice(txRows, func(t *ent.Txn) uint64 { return t.ID })
	_productRows, err := productService.ListInIDs(ctx, _ids)
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}
	// map(ID => *model.WalletProduct)
	productIDMap := util.KeyFunc(_productRows, func(c *ent.Product) int64 {
		return int64(c.ID)
	})

	// 获取区块配置
	bCfgs, err := s.ioc.BlockchainService().GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}
	bcMaps := util.KeyFunc(bCfgs, func(e *ent.Blockchain) uint64 {
		return e.ID
	})

	// 通知id集合
	var notifyIDs []uint64
	// 通知数据
	var notifyRows []*ent.Notify
	// 处理通知数据
	for _, txRow := range txRows {
		productRow, ok := productIDMap[txRow.ProductID]
		if !ok {
			log.Error().Err(err).Msgf("no productMap: %d", txRow.ProductID)
			notifyIDs = append(notifyIDs, txRow.ID)
			continue
		}

		// 构建通知post提交数据
		bc := bcMaps[txRow.ChainID]
		bodyObj := gin.H{
			"appid":   productRow.AppID,
			"address": txRow.ToAddress,
			"chain":   bc.Chain,
			// "type":        bc.Type,
			"symbol":      bc.Symbol,
			"amount":      txRow.AmountStr,
			"tx_hash":     txRow.TxID,
			"notify_type": notify.NotifyTypeReceived,
		}
		bodyObj["sign"] = util.WechatSign(productRow.AppSecret, bodyObj)
		body, err := json.Marshal(bodyObj)
		if err != nil {
			log.Error().Err(err).Msg("err")
			continue
		}
		notifyIDs = append(notifyIDs, txRow.ID)
		notifyRows = append(notifyRows, &ent.Notify{
			Nonce:        uuid.New().String(),
			ChainID:      1,
			ProductID:    uint64(txRow.ProductID),
			ItemType:     1, // 充值
			ItemFrom:     uint64(txRow.ID),
			NotifyType:   notify.NotifyTypeReceived,
			SendURL:      productRow.WebHook,
			SendBody:     string(body),
			HandleStatus: notify.HandleStatusInit.Int(),
			HandleMsg:    notify.HandleStatusInit.String(),
			HandleTime:   time.Now(),
		})
	}
	// 添加通知数据
	if len(notifyRows) > 0 {
		fmt.Println("notifyRows", notifyRows)
		_, err := notifyService.CreateMany(ctx, notifyRows)
		if err != nil {
			log.Error().Err(err).Msg("添加通知数据")
			return
		}
	}
	// 更新交易处理状态
	if len(notifyIDs) > 0 {
		_, err := txnService.UpdateNotifyStatusByIds(ctx, notifyIDs)
		if err != nil {
			log.Error().Err(err).Msg("err")
			return
		}
	}
}

// 检测发送回调通知
func (s *Service) CheckWebhook() {
	log := s.logger
	ctx := context.Background()
	notifyService := s.ioc.NotifyService()
	// 获取待发送通知的数据（包含发送失败的）
	notifyRows, err := notifyService.ListByHandleStatus(ctx, []notify.Handle{
		notify.HandleStatusInit,
		notify.HandleStatusFail,
	})
	if err != nil {
		log.Error().Err(err).Msg("err")
		return
	}
	if len(notifyRows) <= 0 {
		log.Info().Err(err).Msg("暂无处理数据")
		return
	}

	// 遍历通知数据,处理通知操作
	for _, notifyRow := range notifyRows {
		if notifyRow.SendURL == "" {
			continue
		}
		item := notifyRow
		res, err := webhook.HttpRequest(item.SendURL, item.SendBody)
		if err != nil {
			log.Warn().Err(err).Msg("HttpRequest")
			item.HandleMsg = res
			item.SendRetry += 1
			item.HandleStatus = notify.HandleStatusFail.Int()
			_, err := notifyService.UpdateByID(ctx, item)
			if err != nil {
				log.Error().Err(err).Msg("UpdateByID")
				continue
			}
		} else {
			item.HandleMsg = res
			item.HandleStatus = notify.HandleStatusSuccess.Int()
			_, err = notifyService.UpdateByID(ctx, item)
			if err != nil {
				log.Error().Err(err).Msg("UpdateByID")
				continue
			}
			log.Info().
				Str("id", fmt.Sprint(notifyRow.ID)).
				Str("item-type", fmt.Sprint(notifyRow.ItemType)).
				Str("notify-type", notifyRow.NotifyType).
				Msg("通知已发送")
		}
	}
}
