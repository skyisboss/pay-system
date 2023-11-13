package apprun

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/ent/apprun"
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

func (s *Service) Lock(ctx context.Context, handler string) bool {
	s.logger.Info().Str("handler", handler).Msg("开始执行")
	ret, err := s.client.Apprun.Query().Where(apprun.HandlerEQ(handler)).First(ctx)
	if err != nil {

		// 数据为空时新建
		if ent.IsNotFound(err) {
			return s.Add(ctx, handler)
		} else {
			s.logger.Error().Err(err).Str("handler", handler).Msg("Query")
			return false
		}

	}

	// 如果锁已被释放 重新上锁
	if ret.Runing == 0 {
		return s.UpdateLock(ctx, ret.Handler, 1)
	}

	// 如果发生超时死锁 则更新上锁状态
	if time.Now().Unix()-ret.CreatedAt.Unix() > 60*30 {
		return s.UpdateLock(ctx, ret.Handler, 1)
	}

	return false
}

func (s *Service) UnLock(ctx context.Context, handler string) bool {
	s.logger.Info().Str("handler", handler).Msg("执行完毕")
	return s.UpdateLock(ctx, handler, 0)
}

func (s *Service) Add(ctx context.Context, handler string) bool {
	_, err := s.client.Apprun.Create().
		SetCreatedAt(time.Now()).
		SetHandler(handler).
		SetRuning(0).
		SetTotal(1).
		Save(ctx)
	if err != nil {
		s.logger.Error().Err(err).Str("handler", handler).Msg("Add")
	}

	return err == nil
}

func (s *Service) UpdateLock(ctx context.Context, handler string, lock uint64) bool {

	ret, err := s.client.Apprun.
		Update().
		Where(apprun.HandlerEQ(handler)).
		SetRuning(lock).
		SetUpdatedAt(time.Now()).
		AddTotal(int64(lock)).
		Save(ctx)
	if err != nil {
		s.logger.Error().Err(err).Str("handler", handler).Str("lock", string(rune(lock))).Msg("UpdateLock")
	}

	return ret > 0 && err == nil
}
