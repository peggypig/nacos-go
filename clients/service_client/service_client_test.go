package service_client

import (
	"github.com/golang/mock/gomock"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/mock"
	"github.com/peggypig/nacos-go/vo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-15 23:25
**/

func TestServiceClient_goBeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs/listener"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"Listening-Configs": "TEST" + constant.SPLIT_CONFIG_INNER + "TEST" + constant.SPLIT_CONFIG_INNER +
				constant.SPLIT_CONFIG_INNER + constant.SPLIT_CONFIG,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)


	// 错误参数
	client.beating = true
	client.goBeat(constant.ClientConfig{
		BeatInterval:   10 * 1000,
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, []constant.ServerConfig{}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "",
	})
	assert.Equal(t, false, client.beating)

	client.beating = true
	client.goBeat(constant.ClientConfig{
		BeatInterval:   10 * 1000,
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, []constant.ServerConfig{}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "",
		Dom: "DEMO",
	})
	assert.Equal(t, false, client.beating)

	client.beating = true
	client.goBeat(constant.ClientConfig{
		BeatInterval:   10 * 1000,
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, []constant.ServerConfig{}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "demo",
	})
	assert.Equal(t, true, client.beating)
}

func TestServiceClient_stopBeatTask(t *testing.T) {
	client := ServiceClient{}
	client.beating = true
	client.StopBeatTask()
	assert.Equal(t, false, client.beating)
}
