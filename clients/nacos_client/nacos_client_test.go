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
* @create : 2019-01-15 11:59
**/

var clientConfig = constant.ClientConfig{
	TimeoutMs:      10 * 1000,
	BeatInterval:   10 * 1000,
	ListenInterval: 10 * 1000,
}

func TestNacosClient_GetClientConfig(t *testing.T) {
	client := NacosClient{}
	_, err := client.GetClientConfig()
	assert.NotNil(t, err)

	_ = client.SetClientConfig(clientConfig)
	config, err := client.GetClientConfig()
	assert.Nil(t, err)
	assert.Equal(t, clientConfig, config)
}

func TestNacosClient_GetServerConfig(t *testing.T) {
	client := NacosClient{}
	_, err := client.GetServerConfig()
	assert.NotNil(t, err)

	_ = client.SetServerConfig(serverConfigs)
	configs, err := client.GetServerConfig()
	assert.Nil(t, err)
	assert.Equal(t, serverConfigs, configs)
}

func TestNacosClient_SetClientConfig(t *testing.T) {
	testNacosClientSetClientConfig(t, constant.ClientConfig{
		TimeoutMs: 0,
	}, false, clientConfig)

	testNacosClientSetClientConfig(t, constant.ClientConfig{
		TimeoutMs:      1,
		ListenInterval: 2,
	}, false, clientConfig)

	testNacosClientSetClientConfig(t, constant.ClientConfig{
		TimeoutMs: 10,
	}, true, constant.ClientConfig{
		TimeoutMs:      10,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	})

	testNacosClientSetClientConfig(t, clientConfig, true, clientConfig)
}

func testNacosClientSetClientConfig(t *testing.T, config constant.ClientConfig, errNil bool, expected constant.ClientConfig) {
	client := NacosClient{}
	err := client.SetClientConfig(config)
	if errNil {
		assert.Nil(t, err)
		assert.Equal(t, expected, client.clientConfig)
	} else {
		assert.NotNil(t, err)
	}
}

var sreverConfig = constant.ServerConfig{
	IpAddr:      "console.nacos.io",
	Port:        80,
	ContextPath: "/nacos",
}

var serverConfigs = []constant.ServerConfig{sreverConfig}

func TestNacosClient_SetServerConfig(t *testing.T) {
	testNacosClientSetServerConfig(t, []constant.ServerConfig{}, false, sreverConfig)

	testNacosClientSetServerConfig(t, []constant.ServerConfig{
		{
			IpAddr: "console.nacos.io",
			Port:   0,
		},
	}, false, sreverConfig)

	testNacosClientSetServerConfig(t, []constant.ServerConfig{
		{
			IpAddr: "",
			Port:   1,
		},
	}, false, sreverConfig)

	testNacosClientSetServerConfig(t, []constant.ServerConfig{
		{
			IpAddr: "console.nacos.io",
			Port:   65536,
		},
	}, false, sreverConfig)



	testNacosClientSetServerConfig(t, []constant.ServerConfig{
		{
			IpAddr: "console.nacos.io",
			Port:   80,
		},
	}, true, sreverConfig)


	testNacosClientSetServerConfig(t, []constant.ServerConfig{sreverConfig}, true, sreverConfig)
}

func testNacosClientSetServerConfig(t *testing.T, configs []constant.ServerConfig, errNil bool, expected constant.ServerConfig) {
	client := NacosClient{}
	err := client.SetServerConfig(configs)
	if errNil {
		assert.Nil(t, err)
		assert.Equal(t, 1, len(client.serverConfigs))
		assert.Equal(t, expected, client.serverConfigs[0])
	} else {
		assert.NotNil(t, err)
	}
}

func TestNacosClient_SetHttpAgent(t *testing.T) {
	client := NacosClient{}
	agent := http_agent.HttpAgent{}
	err := client.SetHttpAgent(&agent)
	assert.Nil(t, err)
	assert.Equal(t, &agent, client.agent)
}

var namespace = vo.Namespace{
	Namespace:         "4a1515fa-4818-482a-bc49-e4b1a729659b",
	NamespaceShowName: "2345",
	Quota:             200,
	ConfigCount:       0,
	Type:              2,
}

func TestMockINacosClient_GetNamespace(t *testing.T) {
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
		Return(http_agent.FakeHttpResponse(200, `{"code":200,"message":null,
"data":[{"namespace":"4a1515fa-4818-482a-bc49-e4b1a729659b","namespaceShowName":"2345","quota":200,"configCount":0,"type":2}]}`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	content, err := client.GetNamespace()
	assert.Equal(t, []vo.Namespace{namespace}, content)
	assert.Equal(t, nil, err)
}

func TestMockINacosClient_CreateNamespace(t *testing.T) {
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
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 正确参数
	content, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)

	// 错误参数
	_, err = client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "",
	})
	assert.NotNil(t, err)

	_, err = client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "",
		NamespaceDesc: "nacos-go",
	})
	assert.NotNil(t, err)
}

func TestMockINacosClient_DeleteNamespace(t *testing.T) {
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
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)
	// 正确参数
	content, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "5394637d-daf4-4d1c-9075-7c5f733005e8",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)

	_, err = client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "",
	})
	assert.NotNil(t, err)
}

func TestMockINacosClient_ModifyNamespace(t *testing.T) {
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
	_ = client.SetClientConfig(clientConfig)
	_ = client.SetServerConfig(serverConfigs)

	// 正确参数
	content, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceDesc: "nacos-go",
		NamespaceName: "go-nacos",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)

	// 错误参数
	_, err = client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceDesc: "nacos-go",
		NamespaceName: "",
	})
	assert.NotNil(t, err)
	_, err = client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceDesc: "",
		NamespaceName: "nacos-go",
	})
	assert.NotNil(t, err)

	_, err = client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "",
		NamespaceDesc: "nacos-go",
		NamespaceName: "nacos-go",
	})
	assert.NotNil(t, err)

}
