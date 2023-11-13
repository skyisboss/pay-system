package cmd

import (
	"context"

	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/spf13/cobra"
)

var jobCommand = &cobra.Command{
	Use:     "task",
	Short:   "启动任务处理器",
	Example: "oxygen ",
	Run:     runTask,
}

func runTask(_ *cobra.Command, _ []string) {
	var (
		ctx  = context.Background()
		cfg  = RegisterConfig()
		boot = boot.New(ctx, cfg)
		// logger = boot.Logger()
	)

	boot.RunServerTask()
	// if err := graceful.WaitShutdown(); err != nil {
	// 	logger.Error().Err(err).Msg("unable to shutdown service gracefully")
	// 	return
	// }

	// logger.Info().Msg("shutdown complete")
}
