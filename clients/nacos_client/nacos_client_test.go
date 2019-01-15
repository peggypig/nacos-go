package nacos_client

import (
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
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

func TestNacosClient_GetNamespace(t *testing.T) {
	client := NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, err := client.GetNamespace()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(namespaces)
	}
}

func TestNacosClient_CreateNamespace(t *testing.T) {
	client := NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, err := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go-test",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(namespaces)
	}
}

func TestNacosClient_DeleteNamespace(t *testing.T) {
	client := NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, err := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "bdea94a5-fc7e-4d5e-86ca-ad35ce4cb969",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(namespaces)
	}
}

func TestNacosClient_ModifyNamespace(t *testing.T) {
	client := NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, err := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(namespaces)
	}
}
