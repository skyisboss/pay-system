package session

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/tsession"
)

type Service struct {
	client *ent.Client
	logger *zerolog.Logger
}

func New(client *ent.Client, logger *zerolog.Logger) *Service {
	log := logger.With().Str("channel", "apprun.service").Logger()

	return &Service{
		client: client,
		logger: &log,
	}
}

func (s *Service) Create(ctx context.Context, name, value, ip string) (*ent.TSession, error) {
	ret, err := s.client.TSession.Create().
		SetKeyName(name).
		SetKeyValue(value).
		SetIP(ip).
		SetCreatedAt(time.Now()).
		Save(ctx)
	return ret, err
}

func (s *Service) GetSession(ctx context.Context, name, value string) (bool, error) {
	ret, err := s.client.TSession.Query().
		Where(tsession.KeyNameEQ(name)).
		Where(tsession.KeyValueEQ(value)).
		Exist(ctx)
	return ret, err
}
