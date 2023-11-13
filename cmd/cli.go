package cmd

import (
	"context"

	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/spf13/cobra"
)

var cliCommand = &cobra.Command{
	Use:     "cli",
	Short:   "启动处理器",
	Example: "oxygen create-user user@gmail.com qwerty123",
	Run:     run,
}

func run(_ *cobra.Command, _ []string) {
	var (
		ctx     = context.Background()
		cfg     = RegisterConfig()
		service = boot.New(ctx, cfg)
		// users   = service.Container()
		logger = service.Logger()
	)

	logger.Info().Msg("info")
	service.Ioc()
}
