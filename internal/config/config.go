package config

import (
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/skyisboss/pay-system/internal/db/msql"
	"github.com/skyisboss/pay-system/internal/db/psql"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
	"github.com/skyisboss/pay-system/internal/log"
	"github.com/skyisboss/pay-system/internal/util"
)

type Config struct {
	// compile-time parameters
	GitCommit     string
	GitVersion    string
	EmbedFrontend bool

	Env       string     `yaml:"env" env:"APP_ENV" env-default:"production" env-description:"Environment [production, local, sandbox]"`
	Debug     bool       `yaml:"debug" env:"APP_DEBUG" env-default:"false" env-description:"Enables debug mode"`
	Logger    log.Config `yaml:"logger"`
	System    System     `yaml:"system"`
	Database  Database   `yaml:"database"`
	Providers Providers  `yaml:"providers"`
	// KMS    KMS    `yaml:"kms"`
	// Notifications Notifications `yaml:"notifications"`
}

type System struct {
	//Server     http.Config       `yaml:"server"`
	//Processing processing.Config `yaml:"processing"`
}

type Providers struct {
	SaltKey string `yaml:"salt_key" env:"SALT_KEY"`
	EthRpc  string `yaml:"eth_rpc"`
	EthGas  string `yaml:"eth_gas"`
	TronRpc string `yaml:"tron_rpc"`
}

type Database struct {
	Postgres psql.Config `yaml:"postgres"`
	Mysql    msql.Config `yaml:"mysql"`
}

var once = sync.Once{}
var cfg = &Config{}
var errCfg error

func New(gitCommit, gitVersion, configPath string, skipConfig, embedFrontend bool) (*Config, error) {
	once.Do(func() {
		cfg = &Config{
			GitCommit:     gitCommit,
			GitVersion:    gitVersion,
			EmbedFrontend: embedFrontend,
		}

		if skipConfig {
			errCfg = cleanenv.ReadEnv(cfg)
			return
		}

		errCfg = cleanenv.ReadConfig(configPath, cfg)
	})

	return cfg, errCfg
}

func PrintUsage(w io.Writer) error {
	desc, err := cleanenv.GetDescription(&Config{}, nil)
	if err != nil {
		return err
	}

	const delimiter = "||"

	// 1 line == 1 env var
	desc = strings.ReplaceAll(desc, "\n    \t", delimiter)

	lines := strings.Split(desc, "\n")

	// remove header
	lines = lines[1:]

	// hide internal vars
	lines = util.FilterSlice(lines, func(line string) bool {
		return !strings.Contains(strings.ToLower(line), "internal variable")
	})

	// remove duplicates
	lines = lo.Uniq(lines)

	// sort a-z (skip header)
	sort.Strings(lines[1:])

	// write as a table
	t := tablewriter.NewWriter(w)
	t.SetBorder(false)
	t.SetAutoWrapText(false)
	t.SetHeader([]string{"ENV", "Description"})
	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	for _, line := range lines {
		cells := strings.Split(line, delimiter)
		cells = util.MapSlice(cells, strings.TrimSpace)
		t.Append(cells)
	}

	t.Render()

	return nil
}
