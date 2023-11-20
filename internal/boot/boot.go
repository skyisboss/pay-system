package boot

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/skyisboss/pay-system/internal/app/api/router"
	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/config"
	"github.com/skyisboss/pay-system/internal/ioc"
	"github.com/skyisboss/pay-system/internal/log"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type Boot struct {
	config    *config.Config
	ctx       context.Context
	logger    *zerolog.Logger
	services  *ioc.Container
	beforeRun []BeforeRun
}
type BeforeRun func(ctx context.Context, b *Boot) error

func New(ctx context.Context, cfg *config.Config) *Boot {
	hostname, _ := os.Hostname()
	logger := log.New(cfg.Logger, "payke", cfg.GitVersion, cfg.Env, hostname)
	services := ioc.New(ctx, cfg, &logger)

	return &Boot{
		config:   cfg,
		ctx:      ctx,
		logger:   &logger,
		services: services,
	}
}

func (b *Boot) OnBeforeRun(fn BeforeRun) {
	b.beforeRun = append(b.beforeRun, fn)
}

func (b *Boot) Ioc() *ioc.Container {
	return b.services
}

func (b *Boot) Logger() *zerolog.Logger {
	return b.logger
}

// 对外api服务
func (b *Boot) RunServerApi() {
	// app := gin.New()
	// app.Use(gin.Logger(), gin.Recovery())
	// router.NewApi(app, b.Ioc())
	app := InitServer(b.Ioc(), "api")
	app.Run(":52088")
}

// 后台服务
func (b *Boot) RunServerAdmin() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	router.NewAdmin(app, b.Ioc())

	app.Run(":52033")
}

// 定时任务
func (b *Boot) RunServerTask() {

	logger := b.logger.With().Str("channel", "task").Logger()
	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON, Container: b.Ioc()})
	tasks := task.New(
		b.Ioc(),
		taskHandler,
		&logger,
	)
	tasks.CheckWebhook()
	// tasks.EthProvider().CheckDeposit()
	// tasks.EthProvider().CheckCollect()
	// tasks.EthProvider().CheckDeposit()
	/*
		c := cron.New(
			cron.WithSeconds(),
			cron.WithChain(
				cron.Recover(cron.DefaultLogger),
			),
		)
		EthProvider := tasks.EthProvider()

		var err error
		if _, err = c.AddFunc("@every 1s", EthProvider.CheckFreeAddress); err != nil {
			fmt.Printf("err: %s", err)
			return
		}
		c.Start()
	*/
}

func InitServer(ioc *ioc.Container, intiRouter string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	if intiRouter == "api" {
		router.NewApi(r, ioc)
	} else {
		router.NewAdmin(r, ioc)
	}

	return r
}
