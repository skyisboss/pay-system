package address

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/addres"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "address_service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

func (s *Service) GetCountByChainID(ctx context.Context, chainID uint64) (int64, error) {
	count, err := s.client.Addres.Query().Where(addres.UseToEQ(0), addres.ChainID(chainID)).Count(ctx)
	return int64(count), err
}

func (s *Service) CreateMany(ctx context.Context, address []*wallet.Wallet) (int64, error) {
	rows := []*ent.AddresCreate{}
	for _, v := range address {
		row := s.client.Addres.Create().
			SetUUID(v.UUID).
			// SetChainID(v.ChainID).
			SetAddress(v.Address).
			SetPassword(v.PrivateKey).
			SetCreatedAt(time.Now()).
			SetUseTo(0)
		rows = append(rows, row)
	}

	as, err := s.client.Addres.CreateBulk(rows...).Save(ctx)
	return int64(len(as)), err
}

// 查找地址是否存在数据库
func (s *Service) ListByAddressIn(ctx context.Context, addressRows []string) (data []*ent.Addres, err error) {
	ret, err := s.client.Addres.
		Query().
		Where(addres.AddressIn(addressRows...)).
		// Where(addres.UseToGTE(0)).
		All(ctx)
	return ret, err
}

func (s *Service) GetByAddress(ctx context.Context, address string) (*ent.Addres, error) {
	ret, err := s.client.Addres.
		Query().
		Where(addres.AddressEQ(address)).
		First(ctx)
	return ret, err
}
