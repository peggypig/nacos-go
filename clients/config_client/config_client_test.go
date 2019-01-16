package config_client

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
* @create : 2019-01-15 23:05
**/

func TestConfigClient_goListen(t *testing.T) {
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

	client := ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	cfg := vo.ConfigParam{
		DataId: "dataId",
		Group:  "group",
	}
	// 错误参数测试
	client.listening = true
	client.goListen(constant.ClientConfig{
		BeatInterval:   10 * 1000,
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, false, client.listening)
	// 正确参数测试
	client.listening = true
	client.localConfigs = []vo.ConfigParam{cfg}
	client.goListen(constant.ClientConfig{
		BeatInterval:   10 * 1000,
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, true, client.listening)
}

func TestConfigClient_StopListenConfig(t *testing.T) {
	client := ConfigClient{}
	client.listening = true
	client.StopListenConfig()
	assert.Equal(t, false, client.listening)
}

func TestConfigClient_updateLocalConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := mock.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		ContextPath: "/nacos",
		Port:        80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{
			"dataId": "TEST",
			"group":  "TEST",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)

	client.localConfigs = []vo.ConfigParam{{
		DataId: "dataId",
		Group:  "group",
	}}
	client.updateLocalConfig("dataId%02group")
	assert.Equal(t, 1, len(client.localConfigs))
	assert.Equal(t, "MOCK RESULT", client.localConfigs[0].Content)

}

func TestConfigClient_sync(t *testing.T) {
	client := ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	cc := constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   10 * 1000,
	}
	_ = client.SetClientConfig(cc)
	sc := constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}
	_ = client.SetServerConfig([]constant.ServerConfig{
		sc,
	})
	ac := http_agent.HttpAgent{}
	_ = client.SetHttpAgent(&ac)
	clientConfig, serverConfigs, agent, err := client.sync()
	assert.Equal(t, cc, clientConfig)
	assert.Equal(t, 1, len(serverConfigs))
	sc.ContextPath = "/nacos"
	assert.Equal(t, sc, serverConfigs[0])
	assert.Equal(t, &ac, agent)
	assert.Equal(t, nil, err)
}

func TestConfigClient_deleteConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := mock.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId": "TEST",
			"group":  "TEST",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(404, `config not found`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})

	// 正确参数
	content, err := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	assert.NotNil(t,  err)
	assert.Equal(t, false, content)
}
