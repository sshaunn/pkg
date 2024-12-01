package client

import (
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/common"
)

type BitgetApiClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func NewBitgetApiClient(apiKey string, secretKey string, passphrase string) *BitgetApiClient {
	client := common.NewClient(apiKey, secretKey, passphrase)
	return &BitgetApiClient{client}
}

func (p *BitgetApiClient) Post(url string, params map[string]string) (string, error) {
	postBody, jsonErr := pkg.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost(url, postBody)
	return resp, err
}

func (p *BitgetApiClient) Get(url string, params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet(url, params)
	return resp, err
}
