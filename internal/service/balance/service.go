package balance

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/balance"
	"github.com/skyisboss/pay-system/ent/schema"
	"github.com/skyisboss/pay-system/internal/util"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "transaction.service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

// 根据币种配置生成新余额账号
func (s *Service) NewAccount(ctx context.Context, productID uint64) ([]*ent.Balance, error) {
	chainRows, err := s.client.Blockchain.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	chainIDs := util.MapSlice(chainRows, func(e *ent.Blockchain) uint64 {
		return e.ID
	})

	rows := []*ent.BalanceCreate{}
	for _, id := range chainIDs {
		row := s.client.Balance.Create().
			SetChainID(id).
			SetProductID(productID).
			SetBalanceAmount(decimal.NewFromFloat32(0)).
			SetBalanceFreeze(decimal.NewFromFloat32(0)).
			SetTotalDeposit(decimal.NewFromFloat32(0)).
			SetTotalWithdraw(decimal.NewFromFloat32(0)).
			SetCountDeposit(0).
			SetCountWithdraw(0).
			SetVersion(1).
			SetChangeLogs([]schema.ChangeLogs{})

		rows = append(rows, row)
	}
	ret, err := s.client.Balance.CreateBulk(rows...).Save(ctx)
	return ret, err
}

func (s *Service) ListByProductID(ctx context.Context, productID uint64) ([]*ent.Balance, error) {
	ret, err := s.client.Balance.Query().Where(balance.ProductIDEQ(productID)).All(ctx)

	return ret, err
}

type UpdateBalance struct {
	ChainID   uint64
	ProductID uint64
	Amount    decimal.Decimal
}

// 事物 更新充值金额
func (s *Service) TxUpdateDepositAmount(ctx context.Context, tx *ent.Tx, data *UpdateBalance) (int, error) {
	ret, err := tx.Balance.Update().
		Where(balance.ChainIDEQ(data.ChainID)).
		Where(balance.ProductIDEQ(data.ProductID)).
		AddBalanceAmount(data.Amount).
		AddTotalDeposit(data.Amount).
		AddCountDeposit(1).
		Save(ctx)
	return ret, err
}
