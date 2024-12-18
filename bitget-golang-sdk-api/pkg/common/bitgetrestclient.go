package common

import (
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/config"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/constants"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
}

func NewClient(apiKey string, secretKey string, passphrase string) *BitgetRestClient {
	return &BitgetRestClient{
		ApiKey:       apiKey,
		ApiSecretKey: secretKey,
		Passphrase:   passphrase,
		BaseUrl:      config.BaseUrl,
		Signer:       new(Signer).Init(secretKey),
		HttpClient: http.Client{
			Timeout: time.Duration(config.TimeoutSecond) * time.Second,
		},
	}
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := pkg.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	if constants.RSA == config.SignType {
		sign = p.Signer.SignByRSA(constants.POST, uri, params, timesStamp)
	}
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	pkg.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := pkg.TimesStamp()
	body := pkg.BuildGetParams(params)
	//fmt.Println(body)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	pkg.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
