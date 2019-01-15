package nacos_client

import (
	"github.com/golang/mock/gomock"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
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
* @create : 2019-01-15 13:07
**/

func TestMockINacosClient_GetNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `{"code":200,"message":null,
"data":[{"namespace":"4a1515fa-4818-482a-bc49-e4b1a729659b","namespaceShowName":"2345","quota":200,"configCount":0,"type":2}]}`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.GetNamespace()
	assert.Equal(t, []vo.Namespace{{
		Namespace:"4a1515fa-4818-482a-bc49-e4b1a729659b",
		NamespaceShowName:"2345",
		Quota:200,
		ConfigCount:0,
		Type:2},
	}, content)
	assert.Equal(t, nil, err)
}

func TestMockINacosClient_CreateNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func TestMockINacosClient_DeleteNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "5394637d-daf4-4d1c-9075-7c5f733005e8",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func TestMockINacosClient_ModifyNamespace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/console/namespaces"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.AssignableToTypeOf(map[string]string{})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)
	client := NacosClient{}
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceDesc: "nacos-go",
		NamespaceName: "go-nacos",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}
