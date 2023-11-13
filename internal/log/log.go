package log

import (
	"github.com/rs/zerolog"
)

type Config struct {
	Level           string `yaml:"level" env:"LOGGER_LEVEL" env-default:"debug" env-description:"Enabled verbose logging"`
	Pretty          bool   `yaml:"pretty" env:"LOGGER_PRETTY" env-default:"false" env-description:"Enables human readable logging. Otherwise, uses json output"`
	SlackWebhookURL string `yaml:"slack_webhook_url" env:"LOGGER_SLACK_WEBHOOK_URL" env-description:"Internal variable"`
}

func New(cfg Config, serviceName, version, env, host string) zerolog.Logger {
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(level)

	out := zerolog.MultiLevelWriter(
		StdoutWriter(cfg.Pretty),
		SlackWriter(cfg.SlackWebhookURL, zerolog.ErrorLevel),
	)
	return zerolog.New(out).
		With().
		Timestamp().
		// Str("service", serviceName).
		// Str("version", version).
		// Str("env", env).
		// Str("host", host).
		Caller().
		Logger()
}
