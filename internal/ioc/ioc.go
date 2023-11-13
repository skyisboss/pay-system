package ioc

import (
	"context"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/config"
	"github.com/skyisboss/pay-system/internal/rpc/ethrpc"
	"github.com/skyisboss/pay-system/internal/rpc/tronrpc"
	"github.com/skyisboss/pay-system/internal/service/address"
	"github.com/skyisboss/pay-system/internal/service/apprun"
	"github.com/skyisboss/pay-system/internal/service/balance"
	"github.com/skyisboss/pay-system/internal/service/blockchain"
	"github.com/skyisboss/pay-system/internal/service/notify"
	"github.com/skyisboss/pay-system/internal/service/product"
	"github.com/skyisboss/pay-system/internal/service/transfer"
	"github.com/skyisboss/pay-system/internal/service/txn"
	"github.com/skyisboss/pay-system/internal/service/user"
	"github.com/skyisboss/pay-system/internal/service/withdraw"
	"github.com/skyisboss/pay-system/internal/wallet"
)

// 单例模式 容器
type Container struct {
	ctx    context.Context
	config *config.Config
	once   map[string]*sync.Once

	logger *zerolog.Logger

	// Database
	// db       *msql.Connection
	dbClient *ent.Client

	// Clients
	ethClient  *ethrpc.Client
	tronClient *tronrpc.Client
	// trxClient *trx.Client
	// bscClient *bsc.Client

	// Services
	walletService     *wallet.Service
	addressService    *address.Service
	userService       *user.Service
	productService    *product.Service
	blockchainService *blockchain.Service
	txnService        *txn.Service
	balanceService    *balance.Service
	notifyService     *notify.Service
	withdrawService   *withdraw.Service
	transferService   *transfer.Service
	apprunService     *apprun.Service
}

func New(ctx context.Context, cfg *config.Config, logger *zerolog.Logger) *Container {
	return &Container{
		config: cfg,
		ctx:    ctx,
		logger: logger,
		once:   make(map[string]*sync.Once, 128),
	}
}

func (c *Container) DBClient() *ent.Client {
	c.init("ent.client", func() {
		drv, err := sql.Open("mysql", c.config.Database.Mysql.DataSource)
		if err != nil {
			return
		}
		// 获取数据库驱动中的sql.DB对象。
		db := drv.DB()
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)
		ent := ent.NewClient(ent.Driver(drv))

		if err := ent.Schema.Create(c.ctx); err != nil {
			c.logger.Fatal().Err(err).Msgf("failed creating schema resources: %v", err)
			return
		}

		c.dbClient = ent
		if c.config.Database.Mysql.DbDebug {
			c.dbClient = ent.Debug()
		}
	})

	return c.dbClient
}

// func (c *Container) DB() *msql.Connection {
// 	c.init("db", func() {
// 		db, err := msql.Open(c.ctx, c.config.Database.Mysql, c.logger)
// 		if err != nil {
// 			c.logger.Fatal().Err(err).Msg("unable to open database")
// 			return
// 		}

// 		c.db = db

// 		graceful.AddCallback(db.Shutdown)
// 	})

// 	return c.db
// }

func (c *Container) Logger() *zerolog.Logger {
	return c.logger
}

func (c *Container) Config() *config.Config {
	return c.config
}

func (c *Container) WalletService() *wallet.Service {
	c.init("service.wallet", func() {
		providers := wallet.NewProvider().
			AddProvider(&wallet.EthProvider{
				Blockchain: wallet.ETH,
				Client:     c.EthClient(),
			}).
			AddProvider(&wallet.TronProvider{
				Blockchain: wallet.TRON,
				Client:     c.TronClient(),
			})
		c.walletService = wallet.New(c.DBClient(), providers, c.logger)
	})
	return c.walletService
}

func (c *Container) ProductService() *product.Service {
	c.init("service.product", func() {
		c.productService = product.New(c.DBClient(), c.logger)
	})

	return c.productService
}
func (c *Container) TxnService() *txn.Service {
	c.init("service.txn", func() {
		c.txnService = txn.New(c.DBClient(), c.logger)
	})

	return c.txnService
}
func (c *Container) BalanceService() *balance.Service {
	c.init("service.balance", func() {
		c.balanceService = balance.New(c.DBClient(), c.logger)
	})

	return c.balanceService
}
func (c *Container) NotifyService() *notify.Service {
	c.init("service.notify", func() {
		c.notifyService = notify.New(c.DBClient(), c.logger)
	})

	return c.notifyService
}
func (c *Container) WithdrawService() *withdraw.Service {
	c.init("service.withdraw", func() {
		c.withdrawService = withdraw.New(c.DBClient(), c.logger)
	})

	return c.withdrawService
}
func (c *Container) TransferService() *transfer.Service {
	c.init("service.transfer", func() {
		c.transferService = transfer.New(c.DBClient(), c.logger)
	})

	return c.transferService
}

func (c *Container) BlockchainService() *blockchain.Service {
	c.init("service.blockchain", func() {
		c.blockchainService = blockchain.New(c.DBClient(), c.logger)
	})

	return c.blockchainService
}

func (c *Container) UserService() *user.Service {
	c.init("service.user", func() {
		c.userService = user.New(c.DBClient(), c.logger)
	})

	return c.userService
}

func (c *Container) AddressService() *address.Service {
	c.init("service.address", func() {
		c.addressService = address.New(c.DBClient(), c.logger)
	})

	return c.addressService
}

func (c *Container) ApprunService() *apprun.Service {
	c.init("service.apprun", func() {
		c.apprunService = apprun.New(c.DBClient(), c.logger)
	})

	return c.apprunService
}

func (c *Container) Context() context.Context {
	return c.ctx
}

func (c *Container) EthClient() *ethrpc.Client {
	c.init("service.ethClient", func() {
		client, err := ethrpc.DialContext(c.ctx, c.config.Providers.EthRpc)
		if err != nil {
			c.logger.Fatal().Msgf("eth client dial error: [%T] %s", err, err.Error())
			return
		}
		c.ethClient = client
	})
	return c.ethClient
}

func (c *Container) TronClient() *tronrpc.Client {
	c.init("service.TronClient", func() {
		rpc_url := c.config.Providers.TronRpc
		client, err := tronrpc.New(rpc_url)
		if err != nil {
			c.logger.Fatal().Msgf("init client error: [%T] %s", err, err.Error())
			return
		}
		c.tronClient = client
	})
	return c.tronClient
}

func (c *Container) init(key string, f func()) {
	if c.once[key] == nil {
		c.once[key] = &sync.Once{}
	}

	c.once[key].Do(f)
}
