package user

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/user"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
	db     *ent.UserClient
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "product_service").Logger()

	return &Service{
		client: client,
		logger: &log,
		db:     client.User,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (*ent.User, error) {
	ret, err := s.db.Query().Where(user.ID(uint64(id))).First(ctx)
	return ret, err
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*ent.User, error) {
	ret, err := s.db.Query().Where(user.Username(username)).First(ctx)
	return ret, err
}
