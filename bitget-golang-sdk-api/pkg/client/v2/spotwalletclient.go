package v2

import (
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/common"
)

type SpotWalletApi struct {
	BitgetRestClient *common.BitgetRestClient
}

func NewSpotWalletApi(apiKey string, secretKey string, passphrase string) *SpotWalletApi {
	c := common.NewClient(apiKey, secretKey, passphrase)
	return &SpotWalletApi{c}
}

func (p *SpotWalletApi) Transfer(params map[string]string) (string, error) {
	postBody, jsonErr := pkg.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/transfer", postBody)
	return resp, err
}

func (p *SpotWalletApi) DepositAddress(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-address", params)
	return resp, err
}

func (p *SpotWalletApi) Withdrawal(params map[string]string) (string, error) {
	postBody, jsonErr := pkg.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/withdrawal", postBody)
	return resp, err
}

func (p *SpotWalletApi) WithdrawalRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/withdrawal-records", params)
	return resp, err
}

func (p *SpotWalletApi) DepositRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-records", params)
	return resp, err
}
