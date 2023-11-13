package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent"
	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

func TestWithdraw(t *testing.T) {
	_, _, boot, logger := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)

	tasks.EthProvider().CheckWithdraw()

	fmt.Println(111)
}

func TestWithdrawAdd(t *testing.T) {

	ctx, _, boot, log := Setup()
	// 添加提款数据
	t.Run("add-eth", func(t *testing.T) {
		data := &ent.Withdraw{
			ProductID:    1,
			SerialID:     util.GetUUID(),
			ToAddress:    "0xd08d018cf018e947086b05566b5ef61939937851",
			ChainID:      1,
			AmountStr:    "0.123",
			AmountRaw:    decimal.NewFromFloat(123000000000000000),
			TxHash:       "",
			HandleStatus: 0,
			HandleMsg:    "",
			HandleTime:   time.Now(),
		}
		_, err := boot.Ioc().WithdrawService().CreateNew(ctx, data)
		if err != nil {
			log.Fatal().Err(err).Msg("err")
		}
	})

	t.Run("add-erc20", func(t *testing.T) {
		data := &ent.Withdraw{
			ProductID:    1,
			SerialID:     util.GetUUID(),
			ToAddress:    "0xd08d018cf018e947086b05566b5ef61939937851",
			ChainID:      2,
			AmountStr:    "12.340000",
			AmountRaw:    decimal.NewFromFloat(12340000),
			TxHash:       "",
			HandleStatus: 0,
			HandleMsg:    "",
			HandleTime:   time.Now(),
		}
		_, err := boot.Ioc().WithdrawService().CreateNew(ctx, data)
		if err != nil {
			log.Fatal().Err(err).Msg("err")
		}
	})

	fmt.Println(3)
}
