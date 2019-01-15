package nacos_client

import (
	"github.com/golang/mock/gomock"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
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
"data":[{"namespace":"","namespaceShowName":"Public","quota":200,"configCount":16,"type":0},
{"namespace":"4a1515fa-4818-482a-bc49-e4b1a729659b","namespaceShowName":"2345","quota":200,"configCount":0,"type":2},
{"namespace":"5394637d-daf4-4d1c-9075-7c5f733005e8","namespaceShowName":"nacos-go","quota":200,"configCount":0,"type":2},
{"namespace":"7ea377be-0b44-4293-a349-b2b2da3758fb","namespaceShowName":"222","quota":200,"configCount":0,"type":2},
{"namespace":"856733f1-667b-4bc5-ad08-a965c78e46a4","namespaceShowName":"test222","quota":200,"configCount":4,"type":2},
{"namespace":"a86234f2-3062-4d38-885d-08a3870632d3","namespaceShowName":"1","quota":200,"configCount":0,"type":2},
{"namespace":"aa82df8a-73c9-4b01-9ab0-0b3ae61bdead","namespaceShowName":"6","quota":200,"configCount":0,"type":2},
{"namespace":"b4bdbd78-f7d5-45ba-908d-560d07de1f9c","namespaceShowName":"1","quota":200,"configCount":0,"type":2},
{"namespace":"bd8d20e5-9bdf-4bf2-be03-37ffa8b2f3f1","namespaceShowName":"2333","quota":200,"configCount":0,"type":2}]}`), nil)
	client := NacosClient{}
	t.Log(client.SetHttpAgent(mockIHttpAgent))
	t.Log(client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	}))
	t.Log(client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))
	content, err := client.GetNamespace()
	t.Log(content, err)
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
	t.Log(client.SetHttpAgent(mockIHttpAgent))
	t.Log(client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	}))
	t.Log(client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))
	content, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go",
	})
	t.Log(content, err)
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
	t.Log(client.SetHttpAgent(mockIHttpAgent))
	t.Log(client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	}))
	t.Log(client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))
	content, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "5394637d-daf4-4d1c-9075-7c5f733005e8",
	})
	t.Log(content, err)
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
	t.Log(client.SetHttpAgent(mockIHttpAgent))
	t.Log(client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      11 * 1000,
		ListenInterval: 10 * 1000,
	}))
	t.Log(client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))
	content, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceDesc: "nacos-go",
		NamespaceName: "go-nacos",
	})
	t.Log(content, err)
}
