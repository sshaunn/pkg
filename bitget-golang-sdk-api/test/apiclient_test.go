package test

import (
	"fmt"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/config"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/client"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/client/v2"
	"testing"
)

func Test_PlaceOrder(t *testing.T) {
	client := v2.NewMixOrderClient(config.ApiKey, config.SecretKey, config.PASSPHRASE)

	params := pkg.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.PlaceOrder(params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_post(t *testing.T) {
	c := client.NewBitgetApiClient(config.ApiKey, config.SecretKey, config.PASSPHRASE)

	params := pkg.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := c.Post("/api/mix/v1/order/placeOrder", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get(t *testing.T) {
	c := client.NewBitgetApiClient(config.ApiKey, config.SecretKey, config.PASSPHRASE)

	params := pkg.NewParams()
	params["productType"] = "umcbl"

	resp, err := c.Get("/api/mix/v1/account/accounts", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_params(t *testing.T) {
	c := client.NewBitgetApiClient(config.ApiKey, config.SecretKey, config.PASSPHRASE)

	params := pkg.NewParams()

	resp, err := c.Get("/api/spot/v1/account/getInfo", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_encode_params(t *testing.T) {
	c := client.NewBitgetApiClient(config.ApiKey, config.SecretKey, config.PASSPHRASE)

	params := pkg.NewParams()
	params["symbol"] = "$AIUSDT"
	params["businessType"] = "spot"

	resp, err := c.Get("/api/v2/common/trade-rate", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
