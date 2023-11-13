package notify

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/notify"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

type NotifyType int64
type Handle int64

func (s NotifyType) Int() int64 {
	return int64(s)
}
func (s NotifyType) String() string {
	var res string

	switch int64(s) {
	case 0:
		res = ""
	case 100:
		res = "deposit received"
	case 200:
		res = "withdraw"
	case 300:
		res = "collect"
	}
	return res
}
func (s Handle) Int() int64 {
	return int64(s)
}
func (s Handle) String() string {
	var res string

	switch int64(s) {
	case 0:
		res = "init"
	case 1:
		res = "fail"
	case 2:
		res = "success"
	}
	return res
}

const (
	NotifyTypeSend     string = "send"
	NotifyTypeConfirm  string = "confirm"
	NotifyTypeReceived string = "received"

	HandleStatusInit    Handle = 0
	HandleStatusFail    Handle = 1
	HandleStatusSuccess Handle = 2
	// 发送失败重试次数
	NotifySendRetryLimit int64 = 50

	// 充值类型
	NotifyItemDeposit NotifyType = 100
	// 提款类型
	NotifyItemWithdraw NotifyType = 200
	// 零钱整理类型
	NotifyItemCollect NotifyType = 300
)

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "transaction.service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

// 创建通知数据
func (s *Service) CreateMany(ctx context.Context, data []*ent.Notify) (int, error) {
	rows := []*ent.NotifyCreate{}
	for _, v := range data {
		row := s.client.Notify.Create().
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

	as, err := s.client.Notify.CreateBulk(rows...).Save(ctx)
	return len(as), err
}

// 事物 创建通知数据
func (s *Service) TxCreateMany(ctx context.Context, tx *ent.Tx, data []*ent.Notify) ([]*ent.Notify, error) {
	rows := []*ent.NotifyCreate{}
	for _, v := range data {
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

	as, err := tx.Notify.CreateBulk(rows...).Save(ctx)
	return as, err
}

// 获取待发送通知数据
func (s *Service) ListByHandleStatus(ctx context.Context, status []Handle) ([]*ent.Notify, error) {
	var statusRows []int64
	for _, v := range status {
		statusRows = append(statusRows, v.Int())
	}
	ret, err := s.client.Notify.Query().
		Where(notify.HandleStatusIn(statusRows...)).
		Where(notify.SendRetryLTE(NotifySendRetryLimit)).
		All(ctx)
	return ret, err
}

// 更新通知状态
func (s *Service) UpdateByID(ctx context.Context, data *ent.Notify) (int, error) {
	ret, err := s.client.Notify.Update().
		Where(notify.IDEQ(data.ID)).
		SetHandleMsg(data.HandleMsg).
		SetHandleStatus(data.HandleStatus).
		SetSendRetry(data.SendRetry).
		SetHandleTime(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	return ret, err
	// HandleMsg:    body,
	// HandleStatus: config.NotifyStatus.Fail,
	// HandleTime:   time.Now(),
	// SendRetry:    notifyRow.SendRetry + 1,
}
