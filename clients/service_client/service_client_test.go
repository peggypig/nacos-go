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

var clientConfig = constant.ClientConfig{
	BeatInterval:   10 * 1000,
	TimeoutMs:      10 * 1000,
	ListenInterval: 10 * 1000,
}

var serverConfig = constant.ServerConfig{
	IpAddr:      "console.nacos.io",
	Port:        80,
	ContextPath: "/nacos",
}

var serverConfigs = []constant.ServerConfig{serverConfig}

func TestServiceClient_beatTask(t *testing.T) {
	// 错误参数
	testServiceClientbeatTask(t, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "",
	}, http_agent.FakeHttpResponse(200, `beating`), false)

	testServiceClientbeatTask(t, vo.BeatTaskParam{
		Ip:  "",
		Dom: "dom",
	}, http_agent.FakeHttpResponse(200, `beating`), false)

	// 正确参数
	testServiceClientbeatTask(t, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "demo",
	}, http_agent.FakeHttpResponse(200, `beating`), true)
}

func testServiceClientbeatTask(t *testing.T, param vo.BeatTaskParam, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	if errNil {
		mockIHttpAgent.EXPECT().Post(
			gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/api/clientBeat"),
			gomock.AssignableToTypeOf(http.Header{}),
			gomock.AssignableToTypeOf(uint64(10*1000)),
			gomock.Eq(map[string]string{
				"dom":  "demo",
				"beat": `{"ip":"10.0.0.1","port":0,"weight":0,"dom":"demo","cluster":"","metadata":null}`,
			})).AnyTimes().
			Return(response, nil)
	}
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	client.beating = true
	client.beatTask(clientConfig, serverConfigs, mockIHttpAgent, param)
}

func TestServiceClient_stopBeatTask(t *testing.T) {
	client := ServiceClient{}
	client.beating = true
	client.StopBeatTask()
	assert.Equal(t, false, client.beating)
}

var serviceMock = vo.Service(vo.Service{Dom: "DEMO",
	CacheMillis: 1000, UseSpecifiedURL: false,
	Hosts: []vo.Host{
		vo.Host{Valid: true, Marked: false, InstanceId: "10.10.10.10-8888-a-DEMO", Port: 0x22b8,
			Ip:     "10.10.10.10",
			Weight: 1, Metadata: map[string]string{}, ClusterName: "",
			ServiceName: "", Enable: false}}, Checksum: "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
	LastRefTime:                                        0x163f2da7aa2, Env: "", Clusters: "",
	Metadata: map[string]string(nil)})
var serviceJson = `{
			"dom": "DEMO",
			"cacheMillis": 1000,
			"useSpecifiedURL": false,
			"hosts": [{
				"valid": true,
				"marked": false,
				"instanceId": "10.10.10.10-8888-a-DEMO",
				"port": 8888,
				"ip": "10.10.10.10",
				"weight": 1.0,
				"metadata": {}
			}],
			"checksum": "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
			"lastRefTime": 1528787794594,
			"env": "",
			"clusters": ""
		}`

func TestMockIServiceClient_GetService(t *testing.T) {
	testMockIServiceClientGetService(t, vo.GetServiceParam{
		ServiceName: "DEMO",
		Clusters:    []string{"a"},
	}, http_agent.FakeHttpResponse(200, serviceJson), true)

	testMockIServiceClientGetService(t, vo.GetServiceParam{
		ServiceName: "",
		Clusters:    []string{"a"},
	}, http_agent.FakeHttpResponse(200, serviceJson), false)
}

func testMockIServiceClientGetService(t *testing.T, param vo.GetServiceParam, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)
	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)
	if errNil {
		mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)
		mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)
		mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)
		mockIHttpAgent.EXPECT().Get(
			gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/list"),
			gomock.AssignableToTypeOf(http.Header{}),
			gomock.AssignableToTypeOf(uint64(10*1000)),
			gomock.Eq(map[string]string{
				"serviceName": "DEMO",
				"clusters":    "a",
				"healthyOnly": "false",
			})).Times(1).
			Return(response, nil)
	}

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	service, err := client.GetService(param)

	if errNil {
		assert.Nil(t, err)
		assert.Equal(t, serviceMock, service)
	} else {
		assert.NotNil(t, err)
	}
}

func TestMockIServiceClient_getService(t *testing.T) {
	testMockIServieClientgetService(t, http_agent.FakeHttpResponse(200, serviceJson), true)
	testMockIServieClientgetService(t, http_agent.FakeHttpResponse(200, ""), false)
	testMockIServieClientgetService(t, http_agent.FakeHttpResponse(401, "no auth"), false)
}

func testMockIServieClientgetService(t *testing.T, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, _ := createMock(ctrl)
	path := "http://console.nacos.io:80/nacos/v1/ns/instance"
	timeout := uint64(10 * 1000)
	param := map[string]string{
		"serviceName": "DEMO",
		"ip":          "10.0.0.10",
		"port":        "80",
		"cluster":     "DEFAULT",
	}
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq(path),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(timeout),
		gomock.Eq(param)).Times(1).
		Return(response, nil)
	service, err := getService(mockIHttpAgent, path, timeout, param)
	if errNil {
		assert.Nil(t, err)
		assert.Equal(t, serviceMock, service)
	} else {
		assert.NotNil(t, err)
	}
}

func createMock(ctrl *gomock.Controller) (*mock.MockIHttpAgent, *mock.MockINacosClient) {
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockINacosClient := mock.NewMockINacosClient(ctrl)
	return mockIHttpAgent, mockINacosClient
}

func TestMockIServiceClient_GetServiceDetail(t *testing.T) {
	testMockIServiceClientGetServiceDetail(t, http_agent.FakeHttpResponse(200, `{
			"service":{
				"name":"DEMO",
				"protectThreshold":0.0,
				"app":null,
				"group":null,
				"healthCheckMode":"client",
				"metadata":{}
			},
			"clusters":[]}`), true)

	testMockIServiceClientGetServiceDetail(t, http_agent.FakeHttpResponse(200, ``), false)

	testMockIServiceClientGetServiceDetail(t, http_agent.FakeHttpResponse(401, `no auth`), false)
}

func testMockIServiceClientGetServiceDetail(t *testing.T, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)
	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)
	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)
	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)
	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
		})).Times(1).
		Return(response, nil)
	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	service, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "DEMO",
	})
	if errNil {
		assert.Equal(t, nil, err)
		assert.Equal(t, vo.ServiceDetail(vo.ServiceDetail{Service: vo.ServiceInfo{App: "",
			Group: "", HealthCheckMode: "client", Metadata: map[string]string{},
			Name: "DEMO", ProtectThreshold: 0, Selector: vo.ServiceSelector{Selector: ""}},
			Clusters: []vo.Cluster{}}), service)
	} else {
		assert.NotNil(t, err)
	}
}

var serviceInstanceJson = `{
		"metadata": {},
		"instanceId": "10.10.10.10-8888-DEFAULT-DEMO",
		"port": 8888,
		"service": "DEMO",
		"healthy": false,
		"ip": "10.10.10.10",
		"clusterName": "DEFAULT",
		"weight": 1.0
	}`

var serviceInstanceMock = vo.ServiceInstance{InstanceId: "10.10.10.10-8888-DEFAULT-DEMO", Ip: "10.10.10.10",
	Port: 0x22b8, Metadata: map[string]string{}, Service: "DEMO", Healthy: false, ClusterName: "DEFAULT", Weight: 1}

func TestMockIServiceClient_GetServiceInstance(t *testing.T) {
	testMockIServiceClientGetServiceInstance(t, vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        80,
	}, http_agent.FakeHttpResponse(200, serviceInstanceJson), true)

	testMockIServiceClientGetServiceInstance(t, vo.GetServiceInstanceParam{
		ServiceName: "",
		Ip:          "10.10.10.10",
		Port:        80,
	}, http_agent.FakeHttpResponse(200, serviceInstanceJson), false)

	testMockIServiceClientGetServiceInstance(t, vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "",
		Port:        80,
	}, http_agent.FakeHttpResponse(200, serviceInstanceJson), false)

	testMockIServiceClientGetServiceInstance(t, vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        0,
	}, http_agent.FakeHttpResponse(200, serviceInstanceJson), false)

	testMockIServiceClientGetServiceInstance(t, vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        65536,
	}, http_agent.FakeHttpResponse(200, serviceInstanceJson), false)
}

func testMockIServiceClientGetServiceInstance(t *testing.T, param vo.GetServiceInstanceParam, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)
	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)
	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)

	if errNil {
		mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)
		mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)
		mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)
		mockIHttpAgent.EXPECT().Get(
			gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
			gomock.AssignableToTypeOf(http.Header{}),
			gomock.AssignableToTypeOf(uint64(10*1000)),
			gomock.Eq(
				map[string]string{
					"serviceName": "DEMO",
					"ip":          "10.10.10.10",
					"port":        "80",
					"healthyOnly": "false",
				},
			)).Times(1).
			Return(response, nil)
	}

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	service, err := client.GetServiceInstance(param)
	if errNil {
		assert.Nil(t, err)
		assert.Equal(t, serviceInstanceMock, service)
	} else {
		assert.NotNil(t, err)
	}

}

func TestMockIServiceClient_RegisterServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	success, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func TestMockIServiceClient_ModifyServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/update"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"cluster":     "DEFAULT",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	success, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func TestMockIServiceClient_LogoutServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent, mockINacosClient := createMock(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(clientConfig)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq(serverConfigs)).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(clientConfig, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return(serverConfigs, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.0.0.10",
				"port":        "80",
				"cluster":     "DEFAULT",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	success, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func TestServiceClient_beat(t *testing.T) {
	testServiceClientBeat(t, http_agent.FakeHttpResponse(200, `beating`), true)

	testServiceClientBeat(t, http_agent.FakeHttpResponse(403, `no auth`), false)
}

func testServiceClientBeat(t *testing.T, response *http.Response, errNil bool) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	beatParam := map[string]string{
		"Listening-Configs": "TEST" + constant.SPLIT_CONFIG_INNER + "TEST" + constant.SPLIT_CONFIG_INNER +
			constant.SPLIT_CONFIG_INNER + constant.SPLIT_CONFIG,
	}
	path := "http://console.nacos.io:80/nacos/v1/cs/configs/listener"
	timeout := uint64(10 * 1000)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq(path),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(timeout),
		gomock.Eq(beatParam)).AnyTimes().
		Return(response, nil)
	err := beat(mockIHttpAgent, path, timeout, beatParam)
	if errNil {
		assert.Nil(t, err)
	} else {
		assert.NotNil(t, err)
	}
}
