package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

func TestTxConfirm(t *testing.T) {
	_, _, boot, logger := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)

	tasks.EthProvider().CheckTxConfirm()
	fmt.Println(22)
}

func TestTxConfirmTron(t *testing.T) {
	_, _, boot, logger := Setup()

	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON, Container: boot.Ioc()})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)

	tasks.TronProvider().CheckTxConfirm()
	fmt.Println(22)
}

func Test_GetTxByTxid(t *testing.T) {
	_, _, boot, _ := Setup()
	c := boot.Ioc().TronClient()

	coin, err := c.GRPC.GetTransactionInfoByID("bfc7a31d06120ecbd6f4833f70a0822a1d7f68ec95ec325ae90fb8ae0d3449ad")
	if err != nil {
		t.Fatal(err)
	}
	token, err := c.GRPC.GetTransactionInfoByID("9eda9b213610484b833d4939afdf50c4f9936472d1d979feed2813dcbc256d97")
	if err != nil {
		t.Fatal(err)
	}

	util.ToJson(coin)
	util.ToJson(token)
	util.Println(hex.EncodeToString(token.ContractAddress))
	// util.Println(common.BytesToHexString(token.ContractAddress))
	toAddress := address.HexToAddress(common.BytesToHexString(token.ContractAddress))
	util.Println(toAddress.String())
}
