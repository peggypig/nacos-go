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
* @create : 2019-01-17 13:28
**/
var clientConfigTest = constant.ClientConfig{
	TimeoutMs:      10 * 1000,
	BeatInterval:   5 * 1000,
	ListenInterval: 10 * 1000,
}

var serverConfigTest = constant.ServerConfig{
	IpAddr:      "console.nacos.io",
	Port:        80,
	ContextPath: "/nacos",
}

var headerTest = map[string][]string{
	"Content-Type": {"application/x-www-form-urlencoded"},
}

var serverConfigsTest = []constant.ServerConfig{serverConfigTest}

var httpAgentTest = mock.MockIHttpAgent{}

func cretateServerClientTest() ServiceClient {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	return client
}

// sync

func Test_SyncWithoutClientConfig(t *testing.T) {
	client := cretateServerClientTest()
	_, _, _, err := client.sync()
	assert.NotNil(t, err)
}

func Test_SyncWithoutServerConfig(t *testing.T) {
	client := cretateServerClientTest()
	_ = client.SetClientConfig(clientConfigTest)
	_, _, _, err := client.sync()
	assert.NotNil(t, err)
}

func Test_SyncWithoutHttpAgent(t *testing.T) {
	client := cretateServerClientTest()
	_ = client.SetServerConfig(serverConfigsTest)
	_ = client.SetClientConfig(clientConfigTest)
	_, _, _, err := client.sync()
	assert.NotNil(t, err)
}

func Test_Sync(t *testing.T) {
	client := cretateServerClientTest()
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig(serverConfigsTest)
	_ = client.SetHttpAgent(&httpAgentTest)
	_, _, _, err := client.sync()
	assert.Nil(t, err)
}

// RegisterServiceInstance

func Test_RegisterServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

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
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	success, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_RegisterServiceInstanceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

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
		Return(http_agent.FakeHttpResponse(200, `false`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_RegisterServiceInstanceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_RegisterServiceInstanceWithoutIp(t *testing.T) {
	client := ServiceClient{}
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_RegisterServiceInstanceWithInvalidPort_0(t *testing.T) {
	client := ServiceClient{}
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        0,
	})
	assert.NotNil(t, err)
}

func Test_RegisterServiceInstanceWithInvalidPort_65536(t *testing.T) {
	client := ServiceClient{}
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        65536,
	})
	assert.NotNil(t, err)
}

func Test_RegisterServiceInstanceWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	_, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	assert.NotNil(t, err)
}

// LogoutServiceInstance

func Test_LogoutServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
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
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	success, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_LogoutServiceInstanceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.0.0.10",
				"port":        "80",
				"cluster":     "DEFAULT",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(200, `false`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.0.0.10",
				"port":        "80",
				"cluster":     "DEFAULT",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithoutIp(t *testing.T) {
	client := ServiceClient{}
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "",
		Port:        80,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithInvalidPort_0(t *testing.T) {
	client := ServiceClient{}
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        0,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithInvalidPort_65536(t *testing.T) {
	client := ServiceClient{}
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        65536,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithoutCluster(t *testing.T) {
	client := ServiceClient{}
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        80,
		Cluster:     "",
	})
	assert.NotNil(t, err)
}

func Test_LogoutServiceInstanceWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	_, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "",
		Ip:          "10.0.0.1",
		Port:        80,
		Cluster:     "cluster",
	})
	assert.NotNil(t, err)
}

// ModifyServiceInstance

func Test_ModifyServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
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
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	success, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_ModifyServiceInstanceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
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
		Return(http_agent.FakeHttpResponse(200, `false`), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
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
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithoutIp(t *testing.T) {
	client := ServiceClient{}
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "",
		Port:        80,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithInvalidPort_0(t *testing.T) {
	client := ServiceClient{}
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        0,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithInvalidPort_65536(t *testing.T) {
	client := ServiceClient{}
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        65536,
		Cluster:     "aa",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithoutCluster(t *testing.T) {
	client := ServiceClient{}
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        80,
		Cluster:     "",
	})
	assert.NotNil(t, err)
}

func Test_ModifyServiceInstanceWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	_, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "",
		Ip:          "10.0.0.1",
		Port:        80,
		Cluster:     "cluster",
	})
	assert.NotNil(t, err)
}

// GetService

var serviceJsonTest = `{
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

var serviceTest = vo.Service(vo.Service{Dom: "DEMO",
	CacheMillis: 1000, UseSpecifiedURL: false,
	Hosts: []vo.Host{
		vo.Host{Valid: true, Marked: false, InstanceId: "10.10.10.10-8888-a-DEMO", Port: 0x22b8,
			Ip:     "10.10.10.10",
			Weight: 1, Metadata: map[string]string{}, ClusterName: "",
			ServiceName: "", Enable: false}}, Checksum: "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
	LastRefTime:                                        0x163f2da7aa2, Env: "", Clusters: "",
	Metadata: map[string]string(nil)})

func Test_GetService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/list"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"clusters":    "a",
			"healthyOnly": "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, serviceJsonTest), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	service, err := client.GetService(vo.GetServiceParam{
		ServiceName: "DEMO",
		Clusters:    []string{"a"},
	})
	assert.Nil(t, err)
	assert.Equal(t, serviceTest, service)
}

func Test_GetServiceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/list"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"clusters":    "a",
			"healthyOnly": "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetService(vo.GetServiceParam{
		ServiceName: "DEMO",
		Clusters:    []string{"a"},
	})
	assert.NotNil(t, err)
}

func Test_GetServiceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/list"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"clusters":    "a",
			"healthyOnly": "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetService(vo.GetServiceParam{
		ServiceName: "DEMO",
		Clusters:    []string{"a"},
	})
	assert.NotNil(t, err)
}

func Test_GetServiceWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_, err := client.GetService(vo.GetServiceParam{
		ServiceName: "",
		Clusters:    []string{"a"},
	})
	assert.NotNil(t, err)
}

// GetServiceInstance

var serviceInstanceJsonTest = `{
		"metadata": {},
		"instanceId": "10.10.10.10-8888-DEFAULT-DEMO",
		"port": 8888,
		"service": "DEMO",
		"healthy": false,
		"ip": "10.10.10.10",
		"clusterName": "DEFAULT",
		"weight": 1.0
	}`

var serviceInstanceTest = vo.ServiceInstance{InstanceId: "10.10.10.10-8888-DEFAULT-DEMO", Ip: "10.10.10.10",
	Port: 0x22b8, Metadata: map[string]string{}, Service: "DEMO", Healthy: false, ClusterName: "DEFAULT", Weight: 1}

func Test_GetServiceInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.10.10.10",
				"port":        "80",
				"healthyOnly": "false",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(200, serviceInstanceJsonTest), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	service, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        80,
	})
	assert.Nil(t, err)
	assert.Equal(t, serviceInstanceTest, service)
}

func Test_GetServiceInstanceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.10.10.10",
				"port":        "80",
				"healthyOnly": "false",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_GetServiceInstanceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(
			map[string]string{
				"serviceName": "DEMO",
				"ip":          "10.10.10.10",
				"port":        "80",
				"healthyOnly": "false",
			},
		)).Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)

	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_GetServiceInstanceWithoutIp(t *testing.T) {
	client := ServiceClient{}
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "",
		Port:        80,
	})
	assert.NotNil(t, err)
}

func Test_GetServiceInstanceWithInvalidPort_0(t *testing.T) {
	client := ServiceClient{}
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        0,
	})
	assert.NotNil(t, err)
}

func Test_GetServiceInstanceWithInvalidPort_65536(t *testing.T) {
	client := ServiceClient{}
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.1",
		Port:        65536,
	})
	assert.NotNil(t, err)
}

func Test_GetServiceInstanceWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	_, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "",
		Ip:          "10.0.0.1",
		Port:        80,
	})
	assert.NotNil(t, err)
}

// GetServiceDetail

var serviceDetailJsonTest = `{
			"service":{
				"name":"DEMO",
				"protectThreshold":0.0,
				"app":null,
				"group":null,
				"healthCheckMode":"client",
				"metadata":{}
			},
			"clusters":[]}`
var serviceDetailTest = vo.ServiceDetail{Service: vo.ServiceInfo{App: "",
	Group: "", HealthCheckMode: "client", Metadata: map[string]string{},
	Name: "DEMO", ProtectThreshold: 0, Selector: vo.ServiceSelector{Selector: ""}},
	Clusters: []vo.Cluster{}}

func Test_GetServiceDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, serviceDetailJsonTest), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	service, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "DEMO",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, serviceDetailTest, service)
}

func Test_GetServiceDetailWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, ``), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "DEMO",
	})
	assert.NotNil(t, err)
}

func Test_GetServiceDetailWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(401, ``), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "DEMO",
	})
	assert.NotNil(t, err)
}

func Test_GetServiceDetailWithoutServiceName(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "",
	})
	assert.NotNil(t, err)
}

// StopBeatTask

func Test_StopBeatTask(t *testing.T) {
	client := ServiceClient{}
	client.beating = true
	client.StopBeatTask()
	assert.True(t, !client.beating)
}

// beatTask

func Test_beatTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "demo",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, serviceDetailJsonTest), nil)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/api/clientBeat"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dom":  "demo",
			"beat": `{"ip":"10.0.0.1","port":0,"weight":0,"dom":"demo","cluster":"","metadata":null}`,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, `living`), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	client.beating = true
	err := client.beatTask(clientConfigTest, []constant.ServerConfig{serverConfigTest}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "demo",
	})
	assert.Nil(t, err)
	assert.True(t, client.beating)
}

func Test_beatTaskWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "demo",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, serviceDetailJsonTest), nil)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/api/clientBeat"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dom":  "demo",
			"beat": `{"ip":"10.0.0.1","port":0,"weight":0,"dom":"demo","cluster":"","metadata":null}`,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	client.beating = true
	err := client.beatTask(clientConfigTest, []constant.ServerConfig{serverConfigTest}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "10.0.0.1",
		Dom: "demo",
	})
	assert.Nil(t, err)
	assert.True(t, client.beating)
}

func Test_beatTaskWithoutIp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	client.beating = true
	err := client.beatTask(clientConfigTest, []constant.ServerConfig{serverConfigTest}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "",
		Dom: "demo",
	})
	assert.NotNil(t, err)
}

func Test_beatTaskWithoutDom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	client.beating = true
	err := client.beatTask(clientConfigTest, []constant.ServerConfig{serverConfigTest}, mockIHttpAgent, vo.BeatTaskParam{
		Ip:  "ip",
		Dom: "",
	})
	assert.NotNil(t, err)
}


func Test_SubscribeWithoutServiceName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	client.beating = true
	err := client.Subscribe(vo.SubscribeParam{
	})
	assert.NotNil(t, err)
}

func Test_SubscribeWithoutCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	client.beating = true
	err := client.Subscribe(vo.SubscribeParam{
		ServiceName:"DEMO",
	})
	assert.NotNil(t, err)
}

func Test_subscribe(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(gomock.AssignableToTypeOf(""),
		gomock.Any(),
		gomock.Any(),
		gomock.AssignableToTypeOf(map[string]string{})).Times(1).Return(
		http_agent.FakeHttpResponse(200, serviceJsonTest), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ , err := subscribe(mockIHttpAgent,"http://console.nacos.io/v1/ns/instance/list",clientConfigTest.TimeoutMs,
		map[string]string{
			"ServiceName":"demo",
		})
	assert.Nil(t, err)
}

func Test_subscribeWithErrorResponse_401(t *testing.T)  {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(gomock.AssignableToTypeOf(""),
		gomock.Any(),
		gomock.Any(),
		gomock.AssignableToTypeOf(map[string]string{})).Times(1).Return(
		http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ , err := subscribe(mockIHttpAgent,"http://console.nacos.io/v1/ns/instance/list",clientConfigTest.TimeoutMs,
		map[string]string{
			"ServiceName":"demo",
		})
	assert.NotNil(t, err)
}
