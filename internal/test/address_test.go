package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/skyisboss/pay-system/internal/app/task"
	"github.com/skyisboss/pay-system/internal/app/task/eth"
	"github.com/skyisboss/pay-system/internal/app/task/tron"
	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

// func TestCheckAddress(t *testing.T) {}

func TestAddressDecode(t *testing.T) {
	_, cfg, _, _ := Setup()

	// taskHandler := task.NewProvider().
	// 	AddProvider(&eth.Provider{Blockchain: kms.ETH}).
	// 	AddProvider(&tron.Provider{Blockchain: kms.TRON})
	// tasks := task.New(
	// 	boot.Ioc(),
	// 	taskHandler,
	// 	logger,
	// )
	// p := tasks.EthProvider()

	wt := wallet.Wallet{
		Address:    "0x06edba1bf11ff69bfaaec69cf6fc614856989272",
		UUID:       "ef981af8f03646ebbf69a80a5243e1b4",
		PrivateKey: "68c291b6fd97deb801faea07517d06a80e04bf8ab40c8562f2687b15371b9d6c",
	}
	wt.EncodePrivateKey(cfg.Providers.SaltKey)
	// boot.Ioc().WalletService().PrivateKeyEncode(&wt)
	util.ToJson(wt)

	wt2 := wallet.Wallet{
		Address:    "0x06edba1bf11ff69bfaaec69cf6fc614856989272",
		UUID:       "ef981af8f03646ebbf69a80a5243e1b4",
		PrivateKey: "ctqc+5+nxLTHeY0VRTNsIXASzmUxoI6ilij0zOQ10ES9jzUyHL/knKNKrsoULr0h5pVAFFibcWfWcDZwOiwNm+7gXfZGtqRZi88wC7QUwfM=",
	}
	// boot.Ioc().WalletService().PrivateKeyDecode(&wt2)
	wt.EncodePrivateKey(cfg.Providers.SaltKey)
	util.ToJson(wt2)

	fmt.Println(hex.EncodeToString([]byte("68c291b6fd97deb801faea07517d06a80e04bf8ab40c8562f2687b15371b9d6c")))
}

func TestAddressEncode(t *testing.T) {
	_, cfg, _, _ := Setup()
	wt := wallet.Wallet{
		Address:    "0xD08d018cf018e947086B05566B5eF61939937851",
		UUID:       "ef981af8f03646ebbf69a80a5243e1b4",
		PrivateKey: "695fa46e64acd2b884e37618450bbf6c66cc959fd964bd7e372044893ddaa0e6",
	}
	// boot.Ioc().WalletService().PrivateKeyEncode(&wt)
	wt.EncodePrivateKey(cfg.Providers.SaltKey)
	util.ToJson(wt)
}

func TestCheckAddress(t *testing.T) {
	_, _, boot, logger := Setup()
	taskHandler := task.NewProvider().
		AddProvider(&eth.Provider{Blockchain: wallet.ETH}).
		AddProvider(&tron.Provider{Blockchain: wallet.TRON})
	tasks := task.New(
		boot.Ioc(),
		taskHandler,
		logger,
	)
	fmt.Println(1)

	tasks.CheckFreeAddress()
}

func TestAddressEncodeTron(t *testing.T) {
	_, cfg, _, _ := Setup()
	wt := wallet.Wallet{
		Address:    "TMNXXt5vyCiSo8G4ydknZU7mx7rV8gbvvs",
		UUID:       "94c3ff8a8cb347f38bc060b7b6524944",
		PrivateKey: "a5b3300d7cba845e9daa089574d01a7b656aa6aba3ad999e8352282657b9562b",
	}
	// boot.Ioc().WalletService().PrivateKeyEncode(&wt)
	pk, _ := wt.EncodePrivateKey(cfg.Providers.SaltKey)
	wt.PrivateKey = pk
	util.ToJson(wt)

	pk2, _ := wt.DecodePrivateKey(cfg.Providers.SaltKey)
	util.Println(pk2)
}
