package withdraw

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/withdraw"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

type Withdraw int64

const (
	WithdrawStatusInit Withdraw = 0
	WithdrawStatusHex  Withdraw = 1
	WithdrawStatusSend Withdraw = 2
	WithdrawStatusDone Withdraw = 3
)

func (s Withdraw) Int64() int64 {
	return int64(s)
}
func (s Withdraw) String() string {
	var msg string
	switch s {
	case 0:
		msg = "init"
	case 1:
		msg = "hex"
	case 2:
		msg = "send"
	case 3:
		msg = "complete"
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

func (s *Service) CreateNew(ctx context.Context, data *ent.Withdraw) (*ent.Withdraw, error) {
	ret, err := s.client.Withdraw.Create().
		SetChainID(data.ChainID).
		SetSerialID(data.SerialID).
		SetToAddress(data.ToAddress).
		SetProductID(data.ProductID).
		SetAmountRaw(data.AmountRaw).
		SetAmountStr(data.AmountStr).
		SetHandleStatus(data.HandleStatus).
		SetHandleTime(time.Now()).
		Save(ctx)

	return ret, err
}

type QueryList struct {
	// 区块配置id
	ChainIDs []uint64
	// 状态
	HandleStatus Withdraw
}

// 获取待提币数据
func (s *Service) ListWithdrawByHandleStatus(ctx context.Context, chainIDs []uint64) ([]*ent.Withdraw, error) {
	ret, err := s.client.Withdraw.
		Query().
		Where(withdraw.HandleStatusEQ(WithdrawStatusInit.Int64())).
		Where(withdraw.ChainIDIn(chainIDs...)).
		All(ctx)

	return ret, err
}

// 更新处理状态
func (s *Service) UpdateHandleStatusById(ctx context.Context, id uint64, txHash string, status Withdraw) (int, error) {
	ret, err := s.client.Withdraw.
		Update().
		Where(withdraw.ID(id)).
		SetHandleMsg(status.String()).
		SetHandleStatus(status.Int64()).
		SetHandleTime(time.Now()).
		SetTxHash(txHash).
		Save(ctx)

	return ret, err
}

// 根据ids获取列表
func (s *Service) ListByIDsIn(ctx context.Context, ids []uint64) ([]*ent.Withdraw, error) {
	ret, err := s.client.Withdraw.Query().Where(withdraw.IDIn(ids...)).All(ctx)
	return ret, err
}
