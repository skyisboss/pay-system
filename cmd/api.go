package cmd

import (
	"context"

	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/spf13/cobra"
)

var apiCommand = &cobra.Command{
	Use:     "api",
	Short:   "启动Api http",
	Example: "api ",
	Run:     runApi,
}

func runApi(_ *cobra.Command, _ []string) {
	var (
		ctx  = context.Background()
		cfg  = RegisterConfig()
		boot = boot.New(ctx, cfg)
	)

	boot.RunServerApi()
}
