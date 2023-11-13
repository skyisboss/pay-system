package msql

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DataSource string `yaml:"data_source" env:"MYSQL_DSN"`
	DbDebug    bool   `yaml:"db_debug" env:"DB_DEBUG"`
}

type Connection struct {
	Pool   *gorm.DB
	logger *zerolog.Logger
}

func init() {

}

func Open(ctx context.Context, cfg Config, logger *zerolog.Logger) (*Connection, error) {
	db, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		return nil, errors.Wrap(err, "无法连接数据库")
	}

	log := logger.With().Str("channel", "msql").Logger()

	connection := &Connection{
		Pool:   db.WithContext(ctx),
		logger: logger,
	}
	if cfg.DbDebug {
		connection.Pool = connection.Pool.Debug()
	}

	log.Info().
		//Str("db_host", db.ConnConfig.Host).
		//Str("db_name", dbConfig.ConnConfig.Database).
		//Str("db_user", dbConfig.ConnConfig.User).
		//Int32("db_min_connections", dbConfig.MaxConns).
		//Int32("db_max_connections", dbConfig.MinConns).
		Msg("connected to msql")

	return connection, nil
}
func (c *Connection) Instance() *gorm.DB {
	return c.Pool
}

func (c *Connection) Shutdown() error {
	db, _ := c.Pool.DB()
	c.logger.Info().Msg("shutting down mysql connections")
	err := db.Close()
	return err
}
