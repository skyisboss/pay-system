package test

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/cmd"
	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/skyisboss/pay-system/internal/config"
)

func Setup() (context.Context, *config.Config, *boot.Boot, *zerolog.Logger) {
	var (
		ctx    = context.Background()
		cfg    = cmd.RegisterConfig()
		boot   = boot.New(ctx, cfg)
		logger = boot.Logger()
	)

	return ctx, cfg, boot, logger
}
