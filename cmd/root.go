package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/skyisboss/pay-system/internal/config"
	"github.com/spf13/cobra"
)

var (
	Commit        = "none"
	Version       = "1.0.0"
	EmbedFrontend = false
	configPath    = "config/config.yml"
	skipConfig    = false
)

var rootCmd = &cobra.Command{
	Use:   "payke",
	Short: "payke",
	Long:  "payke",
}

func RegisterConfig() *config.Config {
	path, _ := os.Getwd()
	// 是否在test文件夹执行，/Users/max/Desktop/MySpace/system-go/test/cli
	index := strings.Index(path, "internal/test")
	if index >= 0 {
		list := strings.Split(path, "internal/test")
		configPath = list[0] + configPath
	}

	cfg, err := config.New(Commit, Version, configPath, skipConfig, EmbedFrontend)
	if err != nil {
		fmt.Printf("unable to initialize config: %s\n", err.Error())
		os.Exit(1)
	}

	if skipConfig {
		fmt.Println("Skipped file-based configuration, using only ENV")
	}

	return cfg
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func envCommand(_ *cobra.Command, _ []string) {
// 	if err := config.PrintUsage(os.Stdout); err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

func init() {
	RegisterConfig()

	rootCmd.AddCommand(migrateCommand)
	rootCmd.AddCommand(cliCommand)
	rootCmd.AddCommand(jobCommand)

}
