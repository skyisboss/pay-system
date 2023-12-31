package test

import (
	"fmt"
	"testing"

	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/wallet"
)

func TestChecTxSend(t *testing.T) {
	_, _, boot, logger := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)

	tasks.EthProvider().CheckTxSend()
	fmt.Println(1)
}

func TestChecTxSendTron(t *testing.T) {
	_, _, boot, logger := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON, Container: boot.Ioc()})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)

	tasks.TronProvider().CheckTxSend()
	fmt.Println(111)
}
