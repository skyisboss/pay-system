package transfer

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/transfer"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

type Handle int64

const (
	HandleStatusInit    Handle = 0
	HandleStatusSend    Handle = 1
	HandleStatusConfirm Handle = 2

	// 关联操作类型 1xx-零钱整理
	RelatedTypeCollect      int64 = 100
	RelatedTypeCollectErc20 int64 = 101
	RelatedTypeCollectTrc20 int64 = 102

	// 关联操作类型 2xx-用户提币
	RelatedTypeWithdraw      int64 = 200
	RelatedTypeWithdrawErc20 int64 = 201
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
		msg = "send"
	case 2:
		msg = "confirm"
	}
	return msg
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "transfer.service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

func (s *Service) GetNonceByFromAddress(ctx context.Context, address string) (int, error) {
	ret, err := s.client.Transfer.Query().
		Where(transfer.FromAddressEQ(address)).
		Aggregate(ent.Max(transfer.FieldNonce)).
		Int(ctx)
	if err != nil {
		return 0, nil
	}

	return ret, err
}

func (s *Service) CreateMany(ctx context.Context, data []*ent.Transfer) (int64, error) {
	rows := []*ent.TransferCreate{}
	for _, v := range data {
		row := s.client.Transfer.Create().
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

	as, err := s.client.Transfer.CreateBulk(rows...).Save(ctx)
	return int64(len(as)), err
}

// 待发送交易查询参数
type QuerySend struct {
	// 区块配置id
	ChainIDs []uint64
	// 状态
	HandleStatus Handle
}

// 获取待发送交易数据
func (s *Service) ListByHandleStatus(ctx context.Context, q QuerySend) ([]*ent.Transfer, error) {
	ret, err := s.client.Transfer.
		Query().
		Where(transfer.HandleStatusEQ(q.HandleStatus.Int64())).
		Where(transfer.ChainIDIn(q.ChainIDs...)).
		All(ctx)
	if len(ret) <= 0 {
		return nil, errors.New("没有要处理的数据")
	}
	return ret, err
}

// 更新处理状态
func (s *Service) UpdateHandleStatusByIds(ctx context.Context, ids []uint64, status Handle) (int, error) {
	ret, err := s.client.Transfer.
		Update().
		Where(transfer.IDIn(ids...)).
		SetHandleMsg(status.String()).
		SetHandleStatus(status.Int64()).
		SetHandleTime(time.Now()).
		Save(ctx)
	return ret, err
}

// 更新处理状态
func (s *Service) UpdateStatusByTxIDs(ctx context.Context, txs []string, status Handle) (int, error) {
	ret, err := s.client.Transfer.
		Update().
		Where(transfer.TxIDIn(txs...)).
		SetHandleMsg(status.String()).
		SetHandleStatus(status.Int64()).
		SetHandleTime(time.Now()).
		Save(ctx)
	return ret, err
}

// 更新汽油
func (s *Service) UpdateGasByID(ctx context.Context, id uint64, gas, gasPrice int64) (int, error) {
	ret, err := s.client.Transfer.
		Update().
		Where(transfer.IDEQ(id)).
		SetGas(gas).
		SetGasPrice(gasPrice).
		Save(ctx)
	return ret, err
}

// 更新交易确认状态
func (s *Service) UpdateConfirms(ctx context.Context, id uint64, gas, price int64, status Handle) (int, error) {
	ret, err := s.client.Transfer.
		Update().
		Where(transfer.ID(id)).
		SetHandleMsg(status.String()).
		SetHandleStatus(status.Int64()).
		SetHandleTime(time.Now()).
		SetGas(gas).
		SetGasPrice(price).
		Save(ctx)
	return ret, err
}
