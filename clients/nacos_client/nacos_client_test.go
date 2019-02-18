package nacos_client

import (
	"github.com/golang/mock/gomock"
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
* @create : 2019-01-17 11:03
**/

var clientConfigTest = constant.ClientConfig{
	TimeoutMs:      10 * 1000,
	BeatInterval:   5 * 1000,
	ListenInterval: 10 * 1000,
}

var serverConfigTest = constant.ServerConfig{
	ContextPath: "/nacos",
	Port:        80,
	IpAddr:      "console.nacos.io",
}

func createNacosClientTest() (client NacosClient) {
	client = NacosClient{}
	_ = client.SetHttpAgent(&mock.MockIHttpAgent{})
	return client
}

// SetClientConfig
func Test_SetClientConfig(t *testing.T) {
	client := NacosClient{}
	err := client.SetClientConfig(clientConfigTest)
	assert.Nil(t, err)
	config, _ := client.GetClientConfig()
	clientConfigTest.SubscribeInterval = 10 * 1000
	assert.Equal(t, clientConfigTest, config)
}

func Test_SetClientConfigWithoutTimeoutMs(t *testing.T) {
	client := NacosClient{}
	err := client.SetClientConfig(constant.ClientConfig{
		ListenInterval: 10000,
		BeatInterval:   10000,
	})
	assert.NotNil(t, err)
}

func Test_SetClientConfigWithoutTimeoutMsLessListenInterval(t *testing.T) {
	client := NacosClient{}
	err := client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10000,
		ListenInterval: 11000,
		BeatInterval:   10000,
	})
	assert.NotNil(t, err)
}

func Test_SetClientConfigWithoutBeatIntervalAndListenInterval(t *testing.T) {
	client := NacosClient{}
	err := client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 10000,
	})
	assert.Nil(t, err)
	config, _ := client.GetClientConfig()
	clientConfigTest.SubscribeInterval = 10*1000
	assert.Equal(t, clientConfigTest, config)
}

// SetServerConfig

func Test_SetServerConfig(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	assert.Nil(t, err)
	configs, _ := client.GetServerConfig()
	assert.Equal(t, 1, len(configs))
	assert.Equal(t, serverConfigTest, configs[0])
}

func Test_SetServerConfigWithoutContentPath(t *testing.T) {
	client := NacosClient{}
	config := constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}
	err := client.SetServerConfig([]constant.ServerConfig{
		config,
	})
	assert.Nil(t, err)
	configs, _ := client.GetServerConfig()
	assert.Equal(t, 1, len(configs))
	config.ContextPath = "/nacos"
	assert.Equal(t, serverConfigTest, configs[0])
}

func Test_SetServerConfigWithoutIpAddr(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{
		{
			Port:        80,
			ContextPath: "/nacos",
		},
	})
	assert.NotNil(t, err)
}

func Test_SetServerConfigWithoutPort(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{
		{
			IpAddr:      "console.nacos.io",
			ContextPath: "/nacos",
		},
	})
	assert.NotNil(t, err)
}

func Test_SetServerConfigWithInvalidPort_0(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{
		{
			IpAddr:      "console.nacos.io",
			Port:        0,
			ContextPath: "/nacos",
		},
	})
	assert.NotNil(t, err)
}

func Test_SetServerConfigWithInvalidPort_65536(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{
		{
			IpAddr:      "console.nacos.io",
			Port:        65536,
			ContextPath: "/nacos",
		},
	})
	assert.NotNil(t, err)
}

func Test_SetServerConfigWithoutConfig(t *testing.T) {
	client := NacosClient{}
	err := client.SetServerConfig([]constant.ServerConfig{})
	assert.NotNil(t, err)
}

// GetClientConfig
func Test_GetClientConfig_WithoutSet(t *testing.T) {
	client := NacosClient{}
	_, err := client.GetClientConfig()
	assert.NotNil(t, err)
}

func Test_GetClientConfig(t *testing.T) {
	client := NacosClient{}
	client.clientConfigValid = true
	_, err := client.GetClientConfig()
	assert.Nil(t, err)
}

// GetServerConfig

func Test_GetServerConfig(t *testing.T) {
	client := NacosClient{}
	client.serverConfigsValid = true
	_, err := client.GetServerConfig()
	assert.Nil(t, err)
}

func Test_GetServerConfigWithoutSet(t *testing.T) {
	client := NacosClient{}
	_, err := client.GetServerConfig()
	assert.NotNil(t, err)
}

// SetHttpAgent

func Test_SetHttpAgentWithNil(t *testing.T) {
	client := NacosClient{}
	err := client.SetHttpAgent(nil)
	assert.NotNil(t, err)
}

func Test_SetHttpAgent(t *testing.T) {
	client := NacosClient{}
	err := client.SetHttpAgent(&http_agent.HttpAgent{})
	assert.Nil(t, err)
}

// check

func Test_chec(t *testing.T) {
	client := NacosClient{}
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	err := client.check()
	assert.Nil(t, err)
}

func Test_checkWithoutClientConfig(t *testing.T) {
	client := NacosClient{}
	err := client.check()
	assert.NotNil(t, err)
}

func Test_checkWithoutServerConfig(t *testing.T) {
	client := NacosClient{}
	_ = client.SetClientConfig(clientConfigTest)
	//_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	err := client.check()
	assert.NotNil(t, err)
}

func Test_checkWithoutHttpAgent(t *testing.T) {
	client := NacosClient{}
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	err := client.check()
	assert.NotNil(t, err)
}

var namespaceTest = vo.Namespace{
	Namespace:         "aaa",
	NamespaceShowName: "2345",
	Quota:             200,
	ConfigCount:       0,
	Type:              2,
}

var namespaceResponseTest = `{"code":200,"message":null,
"data":[{"namespace":"aaa","namespaceShowName":"2345","quota":200,"configCount":0,"type":2}]}`

// GetNamespace
func Test_GetNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, namespaceResponseTest), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	namespaces, err := client.GetNamespace()
	assert.Equal(t, []vo.Namespace{namespaceTest}, namespaces)
	assert.Equal(t, nil, err)
}

func Test_GetNamespaceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, ``), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetNamespace()
	assert.NotNil(t, err)
}

func Test_GetNamespaceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.GetNamespace()
	assert.NotNil(t, err)
}

// CreateNamespace

func Test_CreateNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	content, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "name",
		NamespaceDesc: "desc",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func Test_CreateNamespaceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `false`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	_, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "name",
		NamespaceDesc: "desc",
	})
	assert.NotNil(t, err)
}

func Test_CreateNamespaceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.Eq(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	_, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "name",
		NamespaceDesc: "desc",
	})
	assert.NotNil(t, err)
}

func Test_CreateNamespaceWithoutNamespaceName(t *testing.T) {
	client := NacosClient{}
	_, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "",
		NamespaceDesc: "desc",
	})
	assert.NotNil(t, err)
}

func Test_CreateNamespaceWithoutNamespaceDesc(t *testing.T) {
	client := NacosClient{}
	_, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "name",
		NamespaceDesc: "",
	})
	assert.NotNil(t, err)
}

// ModifyNamespace

func Test_ModifyNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	// 正确参数
	success, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "aaa",
		NamespaceDesc: "desc",
		NamespaceName: "name",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_ModifyNamespaceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `false`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	_, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "aaa",
		NamespaceDesc: "desc",
		NamespaceName: "name",
	})
	assert.NotNil(t, err)
}

func Test_ModifyNamespaceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})

	// 正确参数
	_, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "aaa",
		NamespaceDesc: "desc",
		NamespaceName: "name",
	})
	assert.NotNil(t, err)
}

func Test_ModifyNamespaceWithoutNamespace(t *testing.T) {
	client := NacosClient{}
	_, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "",
		NamespaceDesc: "nacos-go",
		NamespaceName: "name",
	})
	assert.NotNil(t, err)
}

func Test_ModifyNamespaceWithoutNamespaceName(t *testing.T) {
	client := NacosClient{}
	_, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "aaa",
		NamespaceDesc: "desc",
		NamespaceName: "",
	})
	assert.NotNil(t, err)
}

func Test_ModifyNamespaceWithoutNamespaceDesc(t *testing.T) {
	client := NacosClient{}
	_, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "aaa",
		NamespaceDesc: "",
		NamespaceName: "name",
	})
	assert.NotNil(t, err)
}

// DeleteNamespace

func Test_DeleteNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	success, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "aaa",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, success)
}

func Test_DeleteNamespaceWithErrorResponse_Body(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `false`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "aaa",
	})
	assert.NotNil(t, err)
}

func Test_DeleteNamespaceWithErrorResponse_401(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := mock.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(401, `no auth`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfigTest)
	_ = client.SetServerConfig([]constant.ServerConfig{serverConfigTest})
	_, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "aaa",
	})
	assert.NotNil(t, err)
}

func Test_DeleteNamespaceWithoutNamspaceId(t *testing.T) {
	client := NacosClient{}
	_, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "",
	})
	assert.NotNil(t, err)
}
