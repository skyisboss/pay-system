package cmd

import (
	"context"

	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/spf13/cobra"
)

var adminCommand = &cobra.Command{
	Use:     "admin",
	Short:   "启动Admin http",
	Example: "api ",
	Run:     runAdmin,
}

func runAdmin(_ *cobra.Command, _ []string) {
	var (
		ctx  = context.Background()
		cfg  = RegisterConfig()
		boot = boot.New(ctx, cfg)
	)

	boot.RunServerAdmin()
}
