package product

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/product"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "product_service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

func (s *Service) GetByID(ctx context.Context, id uint64) (*ent.Product, error) {
	ret, err := s.client.Product.Query().Where(product.ID(id)).First(ctx)
	return ret, err
}

func (s *Service) GetByAppID(ctx context.Context, appid string) (*ent.Product, error) {
	ret, err := s.client.Product.Query().Where(product.AppID(appid)).First(ctx)
	return ret, err
}

func (s *Service) ListByAppName(ctx context.Context, appName string) ([]*ent.Product, error) {
	ret, err := s.client.Product.Query().Where(product.AppName(appName)).All(ctx)
	return ret, err
}

func (s *Service) ListInIDs(ctx context.Context, ids []uint64) ([]*ent.Product, error) {
	ret, err := s.client.Product.Query().Where(product.IDIn(ids...)).All(ctx)
	return ret, err
}

func (s *Service) Create(
	ctx context.Context,
	appID, appName, appSecret, webHook string,
	appStatus, withdrawStatus int64,
) (*ent.Product, error) {
	ret, err := s.client.Product.Create().
		SetAppID(appID).
		SetAppName(appName).
		SetAppSecret(appSecret).
		SetAppStatus(appStatus).
		SetWebHook(webHook).
		SetWithdrawStatus(withdrawStatus).
		SetCreatedAt(time.Now()).
		Save(ctx)
	return ret, err
}
