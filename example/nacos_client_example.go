package example

import (
	"fmt"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-15 11:59
**/


func ExampleNacosClient_GetNamespace() {
	client := nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, _ := client.GetNamespace()
	fmt.Println(namespaces)
}

func ExampleNacosClient_CreateNamespace() {
	client := nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, _ := client.CreateNamespace(vo.CreateNamespaceParam{
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go-test",
	})
	fmt.Println(namespaces)
}

func ExampleNacosClient_DeleteNamespace() {
	client := nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, _ := client.DeleteNamespace(vo.DeleteNamespaceParam{
		NamespaceId: "bdea94a5-fc7e-4d5e-86ca-ad35ce4cb969",
	})
	fmt.Println(namespaces)
}

func ExampleNacosClient_ModifyNamespace() {
	client := nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	namespaces, _ := client.ModifyNamespace(vo.ModifyNamespaceParam{
		Namespace:     "5394637d-daf4-4d1c-9075-7c5f733005e8",
		NamespaceName: "nacos-go",
		NamespaceDesc: "nacos-go",
	})
	fmt.Println(namespaces)
}
