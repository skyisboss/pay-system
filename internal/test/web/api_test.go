package web

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/skyisboss/pay-system/internal/boot"
	"github.com/skyisboss/pay-system/internal/util"
)

func TestWithdraw(t *testing.T) {
	_, _, b, _ := Setup()

	router := boot.InitServer(b.Ioc(), "api")
	jsonData, _ := json.Marshal(gin.H{
		"appid":   "123",
		"nonce":   "18ac31d524b14b759d0de622f21e5a8c",
		"chain":   "eth",
		"symbol":  "eth",
		"address": "0xd08d018cf018e947086b05566b5ef61939937851",
		"amount":  "0.02",
		"sign":    "89606F6AAB6711ECF72383E3E1749FE1",
	})

	w := UseRequest(router, "POST", "/api/v1/withdraw", bytes.NewBuffer(jsonData))
	var res RespType
	json.Unmarshal(w.Body.Bytes(), &res)

	util.ToJson(res)
}
