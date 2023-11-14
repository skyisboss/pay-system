package test

import (
	"testing"

	"github.com/skyisboss/pay-system/internal/util"
	"github.com/skyisboss/pay-system/internal/wallet"
)

func TestTronRpc(t *testing.T) {
	_, _, boot, _ := Setup()
	Provider := boot.Ioc().WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)
	ret, err := Provider.RpcBlockNumber()
	if err != nil {
		util.Println(err)
		util.Println(1)
		return
	}
	util.Println(ret)
}

func TestTronAddress(t *testing.T) {
	_, _, boot, _ := Setup()
	Provider := boot.Ioc().WalletService().GetProvider(wallet.TRON).(*wallet.TronProvider)
	wallet := Provider.CreateWallet()
	wallet.PrivateKey, _ = wallet.EncodePrivateKey("123222")
	util.ToJson(wallet)
}
