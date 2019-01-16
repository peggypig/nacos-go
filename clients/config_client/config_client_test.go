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

var configParam = vo.ConfigParam{
	DataId: "dataId",
	Group:  "group",
}

var clientConfig = constant.ClientConfig{
	BeatInterval:   10 * 1000,
	TimeoutMs:      10 * 1000,
	ListenInterval: 10 * 1000,
}

var serverConfig = constant.ServerConfig{
	IpAddr:      "console.nacos.io",
	ContextPath: "/nacos",
	Port:        80,
}

var paramMap = map[string]string{
	"dataId": "dataId",
	"group":  "group",
}

var serverConfigs = []constant.ServerConfig{serverConfig}

var listenConfig = map[string]string{
	"Listening-Configs": "dataId" + constant.SPLIT_CONFIG_INNER + "group" + constant.SPLIT_CONFIG_INNER +
		constant.SPLIT_CONFIG_INNER + constant.SPLIT_CONFIG,
}

func createMockIHttpAgent(ctrl *gomock.Controller) *mock.MockIHttpAgent {
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	return mockIHttpAgent
}

func TestConfigClient_listenConfigTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := createMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs/listener"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(listenConfig)).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)

	// 正确参数测试
	client.listening = true
	client.localConfigs = []vo.ConfigParam{vo.ConfigParam{
		DataId:  "dataId",
		Group:   "group",
		Tenant:  "aaa",
		Content: "bbb",
	}}
	client.listenConfigTask(clientConfig, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, true, client.listening)

	// 错误参数测试
	client.listening = true
	client.localConfigs = []vo.ConfigParam{}
	client.listenConfigTask(clientConfig, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, false, client.listening)

	client.listening = true
	client.localConfigs = []vo.ConfigParam{vo.ConfigParam{
		DataId: "dataId",
	}}
	client.listenConfigTask(clientConfig, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, false, client.listening)

	client.listening = true
	client.localConfigs = []vo.ConfigParam{vo.ConfigParam{
		Group: "group",
	}}
	client.listenConfigTask(clientConfig, []constant.ServerConfig{}, mockIHttpAgent)
	assert.Equal(t, false, client.listening)
}

func TestConfigClient_listen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := createMockIHttpAgent(ctrl)
	path := "http://console.nacos.io:80/nacos/v1/cs/configs/listener"
	timeout := uint64(10 * 1000)
	listenInterval := timeout

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq(path),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(timeout),
		gomock.Eq(listenConfig)).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)

	// 正确参数测试
	client.localConfigs = []vo.ConfigParam{configParam}
	changed, err := listen(mockIHttpAgent, path, timeout, listenInterval, listenConfig)
	assert.Equal(t, "", changed)
	assert.Nil(t, err)
}

func TestConfigClient_StopListenConfig(t *testing.T) {
	client := ConfigClient{}
	client.listening = true
	client.StopListenConfig()
	assert.Equal(t, false, client.listening)
}

func createMockINacosClient(ctrl *gomock.Controller) *mock.MockINacosClient {
	mockINacosClient := mock.NewMockINacosClient(ctrl)
	return mockINacosClient
}

func TestConfigClient_updateLocalConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := createMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(paramMap)).
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
	cc := clientConfig
	_ = client.SetClientConfig(cc)
	sc := serverConfig
	sc.ContextPath = ""
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
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(paramMap)).
		Times(1).
		Return(http_agent.FakeHttpResponse(404, `config not found`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 正确参数
	content, err := client.DeleteConfig(configParam)
	assert.NotNil(t, err)
	assert.Equal(t, false, content)
}

func TestMockIConfigClient_GetConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(paramMap)).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 正确参数测试
	content, err := client.GetConfig(configParam)
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)

	// 错误参数测试
	_, err = client.GetConfig(vo.ConfigParam{
		DataId: "Test",
	})
	assert.NotNil(t, err)

	_, err = client.GetConfig(vo.ConfigParam{
		Group: "Test",
	})
	assert.NotNil(t, err)
}

func TestMockIConfigClient_getConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	path := "http://console.nacos.io:80/nacos/v1/cs/configs"
	timeout := 10000
	mockIHttpAgent, _ := createMock(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq(path),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(timeout)),
		gomock.Eq(paramMap)).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	content, err := getConfig(mockIHttpAgent, path, uint64(timeout), paramMap)
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)
}

func TestMockIConfigClient_PublishConfig(t *testing.T) {
	testMockIConfigClientPublishConfig(t, http_agent.FakeHttpResponse(200, `true`), true, true)
}

func TestMockIConfigClient_PublishConfig_ResponseBodyException(t *testing.T) {
	testMockIConfigClientPublishConfig(t, http_agent.FakeHttpResponse(200, `false`), false, false)
}

func TestMockIConfigClient_PublishConfig_ResponseStatusException(t *testing.T) {
	testMockIConfigClientPublishConfig(t, http_agent.FakeHttpResponse(403, `no auth`), false, false)
}

func testMockIConfigClientPublishConfig(t *testing.T, response *http.Response, errNil bool, success bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId":  "TEST",
			"group":   "TEST",
			"content": "test",
		})).
		Times(1).
		Return(response, nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 正确参数
	content, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "TEST",
		Content: "test",
	})
	if errNil {
		assert.Nil(t, err)
	} else {
		assert.NotNil(t, err)
	}
	assert.Equal(t, success, content)

	// 错误参数
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "TEST",
		Content: "",
	})
	assert.NotNil(t, err)
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "",
		Content: "TEST",
	})
	assert.NotNil(t, err)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "",
		Group:   "TEST",
		Content: "TEST",
	})
	assert.NotNil(t, err)
}

func createMock(ctrl *gomock.Controller) (*mock.MockIHttpAgent, *mock.MockINacosClient) {
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := mock.NewMockINacosClient(ctrl)
	return mockIHttpAgent, mockINacosClient
}

func TestMockIConfigClient_DeleteConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(paramMap)).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 错误参数
	_, err := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "",
	})
	assert.NotNil(t, err)
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "",
		Group:  "TEST",
	})
	assert.NotNil(t, err)

	// 正确参数
	content, err := client.DeleteConfig(configParam)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func TestMockIConfigClient_GetConfigContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := mock.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(paramMap)).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	// 错误参数
	_, err := client.GetConfigContent("dataId", "")
	assert.NotNil(t, err)
	_, err = client.GetConfigContent("", "dataId")
	assert.NotNil(t, err)

	// 正确参数 没有localConfigs
	content, err := client.GetConfigContent("dataId", "group")
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)

	// 正确参数 有localConfigs
	client.localConfigs = []vo.ConfigParam{
		{DataId: "dataId", Group: "group", Content: "MOCK RESULT"},
	}
	content, err = client.GetConfigContent("dataId", "group")
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)

}
