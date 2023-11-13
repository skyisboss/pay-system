package test

import (
	"fmt"
	"testing"

	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/wallet"
	// "github.com/skyisboss/pay-system/internal/app/kms"
)

func TestCheckCollect(t *testing.T) {
	_, _, boot, log := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&eth.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		log,
	)
	p := tasks.EthProvider()
	p.CheckCollect()
}

func TestCheckCollectErc20(t *testing.T) {
	_, _, boot, log := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		log,
	)
	p := tasks.EthProvider()
	p.CheckCollectErc20()
	fmt.Println(211)
}
