package txn

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/txn"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

type Handle int64
type Collect int64

const (
	HandleStatusInit   Handle = 0
	HandleStatusHex    Handle = 1
	HandleStatusNotify Handle = 1
	HandleStatusSend   Handle = 2

	// 零钱整理状态 - 待整理
	CollectStatusInit Collect = 0
	CollectStatusHex  Collect = 1
	CollectStatusDone Collect = 2

	// 提款状态 - 待整理
	// WithdrawStatusInit Handle = 0
	// WithdrawStatusHex  Handle = 1
	// WithdrawStatusDone Handle = 2
)

func (s Handle) Int64() int64 {
	return int64(s)
}
func (s Handle) String() string {
	var msg string
	switch s {
	case 0:
		msg = "init"
	case 1:
		msg = "notify"
	case 2:
		msg = "send"
	}
	return msg
}
func (s Collect) Int64() int64 {
	return int64(s)
}
func (s Collect) String() string {
	var msg string
	switch s {
	case 1:
		msg = "hex"
	case 2:
		msg = "done"
	default:
		msg = "init"
	}
	return msg
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "transaction.service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

func (s *Service) ListByHandleStatus(ctx context.Context, status int64) ([]*ent.Txn, error) {
	ret, err := s.client.Txn.Query().Where(txn.HandleStatus(status)).All(ctx)
	return ret, err
}

// 零钱整理
// func (s *Service) ListByCollectStatus(ctx context.Context, status int64) ([]*ent.Txn, error) {
// 	ret, err := s.client.Txn.Query().Where(txn.CollectStatusEQ(status)).All(ctx)
// 	return ret, err
// }

// 获取零钱整理列表
func (s *Service) ListByCollectStatusAndChainID(ctx context.Context, status int64, chainID uint64) ([]*ent.Txn, error) {
	ret, err := s.client.Txn.Query().
		Where(txn.CollectStatusEQ(status)).
		Where(txn.ChainIDEQ(chainID)).
		All(ctx)
	return ret, err
}

// 获取零钱整理列表
func (s *Service) ListByCollectChain(ctx context.Context, chainID uint64, amountLimit decimal.Decimal) ([]*ent.Txn, error) {
	ret, err := s.client.Txn.Query().
		Where(txn.ChainIDEQ(chainID)).
		Where(txn.AmountRawGTE(amountLimit)).
		Where(txn.CollectStatusEQ(int64(CollectStatusInit))).
		All(ctx)
	return ret, err
}

// 根据ids批量更新通知数据
func (s *Service) UpdateNotifyStatusByIds(ctx context.Context, ids []uint64) (int, error) {
	ret, err := s.client.Txn.
		Update().
		Where(txn.IDIn(ids...)).
		SetHandleStatus(HandleStatusNotify.Int64()).
		SetHandleMsg("notify").
		SetHandleTime(time.Now()).
		Save(ctx)
	return ret, err
}

// 事物 根据ids批量更新通知数据
func (s *Service) TxUpdateNotifyStatusByIds(ctx context.Context, tx *ent.Tx, ids []uint64) (int, error) {
	ret, err := tx.Txn.
		Update().
		Where(txn.IDIn(ids...)).
		SetHandleStatus(HandleStatusNotify.Int64()).
		SetHandleMsg("notify").
		SetHandleTime(time.Now()).
		Save(ctx)
	return ret, err
}

// 根据ids批量更新零钱整理
func (s *Service) UpdateCollectStatusByIds(ctx context.Context, ids []uint64, status Collect) (int, error) {
	ret, err := s.client.Txn.
		Update().
		Where(txn.IDIn(ids...)).
		SetCollectMsg(status.String()).
		SetCollectStatus(status.Int64()).
		SetCollectTime(time.Now()).
		Save(ctx)
	return ret, err
}

// 根据ids批量更新处理状态
func (s *Service) UpdateHandleStatusByIds(ctx context.Context, ids []uint64, status Handle) (int, error) {
	ret, err := s.client.Txn.
		Update().
		Where(txn.IDIn(ids...)).
		SetHandleMsg(status.String()).
		SetHandleStatus(status.Int64()).
		SetHandleTime(time.Now()).
		Save(ctx)
	return ret, err
}

func (s *Service) CreateMany(ctx context.Context, data []*ent.Txn) ([]*ent.Txn, error) {
	rows := []*ent.TxnCreate{}
	for _, v := range data {
		row := s.client.Txn.Create().
			SetTxID(v.TxID).
			SetChainID(v.ChainID).
			SetProductID(v.ProductID).
			SetFromAddress(v.FromAddress).
			SetToAddress(v.ToAddress).
			SetAmountStr(v.AmountStr).
			SetAmountRaw(v.AmountRaw).
			SetHandleStatus(v.HandleStatus)
		rows = append(rows, row)
	}

	ret, err := s.client.Txn.CreateBulk(rows...).Save(ctx)
	return ret, err
}

func (s *Service) CreateManyWithTx(ctx context.Context, tx *ent.Tx, data []*ent.Txn) ([]*ent.Txn, error) {
	rows := []*ent.TxnCreate{}
	for _, v := range data {
		row := tx.Txn.Create().
			SetTxID(v.TxID).
			SetChainID(v.ChainID).
			SetProductID(v.ProductID).
			SetFromAddress(v.FromAddress).
			SetToAddress(v.ToAddress).
			SetAmountStr(v.AmountStr).
			SetAmountRaw(v.AmountRaw).
			SetHandleStatus(v.HandleStatus)
		rows = append(rows, row)
	}

	ret, err := tx.Txn.CreateBulk(rows...).Save(ctx)
	return ret, err
}

// 根据ids获取列表
func (s *Service) ListByIDsIn(ctx context.Context, ids []uint64) ([]*ent.Txn, error) {
	ret, err := s.client.Txn.Query().Where(txn.IDIn(ids...)).All(ctx)
	return ret, err
}
