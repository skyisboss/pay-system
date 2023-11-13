package test

import (
	"testing"

	"github.com/skyisboss/pay-system/internal/util"
)

func TestTronRpc(t *testing.T) {
	_, _, boot, _ := Setup()

	ret, err := boot.Ioc().TronClient().GRPC.GetNowBlock()
	if err != nil {
		util.Println(err)
		util.Println(1)
		return
	}
	util.Println(ret.BlockHeader.RawData.Number)
}
