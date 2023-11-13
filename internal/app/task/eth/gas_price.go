package eth

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/skyisboss/pay-system/ent/schema"
	"github.com/skyisboss/pay-system/internal/wallet"
)

type GasPriceApi struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		LastBlock       int64   `json:"LastBlock,string"`
		SafeGasPrice    float64 `json:"SafeGasPrice,string"`
		ProposeGasPrice float64 `json:"ProposeGasPrice,string"`
		FastGasPrice    float64 `json:"FastGasPrice,string"`
		SuggestBaseFee  float64 `json:"suggestBaseFee,string"`
		GasUsedRatio    string  `json:"gasUsedRatio"`
	} `json:"result"`
}

// 检测eth汽油价格
func (p *Provider) CheckGasPrice() {
	log := p.ioc.Logger()
	cfg := p.ioc.Config()
	ctx := context.Background()

	gresp, body, errs := gorequest.New().
		Get(cfg.Providers.EthGas).
		Timeout(time.Second * 120).
		End()
	if errs != nil {
		log.Error().Err(errs[0]).Msg("err")
		return
	}
	if gresp.StatusCode != http.StatusOK {
		// 状态错误
		log.Error().Msgf("req status error: %d", gresp.StatusCode)
		return
	}
	var resp GasPriceApi
	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		log.Error().Msgf("err: [%T] %s", err, err.Error())
		return
	}
	if resp.Status != "1" {
		// 状态错误
		log.Error().Msgf("req status error: %s", resp.Status)
		return
	}

	coldGasPrice := int64(1.2 * resp.Result.FastGasPrice * math.Pow10(9))
	userGasPrice := int64(2 * resp.Result.FastGasPrice * math.Pow10(9))
	tipFeePrice := int64(math.Ceil(resp.Result.FastGasPrice-resp.Result.SuggestBaseFee) * math.Pow10(9))
	if tipFeePrice < 0 {
		tipFeePrice = 1 * int64(math.Pow10(9))
	}
	if tipFeePrice > userGasPrice {
		tipFeePrice = coldGasPrice
	}

	// update
	_, err = p.ioc.BlockchainService().UpdateEthGasPrice(ctx, string(wallet.ETH), schema.GasPrice{
		UserGasPrice: userGasPrice,
		ColdGasPrice: coldGasPrice,
		TipFeePrice:  tipFeePrice,
	})
	if err != nil {
		log.Error().Msgf("err: [%T] %s", err, err.Error())
	}
}
