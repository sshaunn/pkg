package ws

import (
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/constants"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/logging/applogger"
	"github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/common"
	model2 "github.com/sshaunn/pkg/bitget-golang-sdk-api/pkg/model"
	"strings"
)

type BitgetWsClient struct {
	bitgetBaseWsClient *common.BitgetBaseWsClient
	NeedLogin          bool
}

func (p *BitgetWsClient) Init(needLogin bool, listener common.OnReceive, errorListener common.OnReceive) *BitgetWsClient {
	p.bitgetBaseWsClient = new(common.BitgetBaseWsClient).Init()
	p.bitgetBaseWsClient.SetListener(listener, errorListener)
	p.bitgetBaseWsClient.ConnectWebSocket()
	p.bitgetBaseWsClient.StartReadLoop()
	p.bitgetBaseWsClient.ExecuterPing()

	if needLogin {
		applogger.Info("login in ...")
		p.bitgetBaseWsClient.Login()
		for {
			if !p.bitgetBaseWsClient.LoginStatus {
				continue
			}
			break
		}
		applogger.Info("login in ... success")
	}

	return p

}

func (p *BitgetWsClient) Connect() *BitgetWsClient {
	p.bitgetBaseWsClient.Connect()
	return p
}

func (p *BitgetWsClient) UnSubscribe(list []model2.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		delete(p.bitgetBaseWsClient.ScribeMap, list[i])
		p.bitgetBaseWsClient.AllSuribe.Add(list[i])
		p.bitgetBaseWsClient.AllSuribe.Remove(list[i])
		args = append(args, list[i])
	}

	wsBaseReq := model2.WsBaseReq{
		Op:   constants.WsOpUnsubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func (p *BitgetWsClient) SubscribeDef(list []model2.SubscribeReq) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := toUpperReq(list[i])
		args = append(args, req)
	}
	wsBaseReq := model2.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.SendMessageByType(wsBaseReq)
}

func toUpperReq(req model2.SubscribeReq) model2.SubscribeReq {
	req.InstType = strings.ToUpper(req.InstType)
	req.InstId = strings.ToUpper(req.InstId)
	req.Channel = strings.ToLower(req.Channel)
	if "" == req.Coin {
		req.Coin = strings.ToLower(req.InstId)
	}
	return req

}

func (p *BitgetWsClient) Subscribe(list []model2.SubscribeReq, listener common.OnReceive) {

	var args []interface{}
	for i := 0; i < len(list); i++ {
		req := toUpperReq(list[i])
		args = append(args, req)

		p.bitgetBaseWsClient.ScribeMap[req] = listener
		p.bitgetBaseWsClient.AllSuribe.Add(req)
		args = append(args, req)
	}

	wsBaseReq := model2.WsBaseReq{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}

	p.bitgetBaseWsClient.SendByType(wsBaseReq)
}

func (p *BitgetWsClient) SendMessage(msg string) {
	p.bitgetBaseWsClient.Send(msg)
}

func (p *BitgetWsClient) SendMessageByType(req model2.WsBaseReq) {
	p.bitgetBaseWsClient.SendByType(req)
}
