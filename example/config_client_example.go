package example

import (
	"fmt"
	"github.com/peggypig/nacos-go/clients/config_client"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 12:05
**/

func ExampleConfigClient_GetConfig() {
	client := config_client.ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "TEST2",
		Group:  "DEFAULT_GROUP",
	})
	fmt.Println(content)
}

func ExampleConfigClient_PublishConfig() {
	client := config_client.ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	content, _ := client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST2",
		Group:   "DEFAULT_GROUP",
		Content: "aaa",
	})
	fmt.Println(content)
}

func ExampleConfigClient_DeleteConfig() {
	client := config_client.ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	content, _ := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
		Tenant: "bbb",
	})
	fmt.Println(content)
}

func ExampleConfigClient_ListenConfig() {
	client := config_client.ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.ListenConfig([]vo.ConfigParam{{
		DataId:  "TEST2",
		Group:   "DEFAULT_GROUP",
		Content: "2019-01-08 09:57:34",
	}, {
		DataId: "test",
		Group:  "DEFAULT_GROUP",
	}})
	time.Sleep(100 * time.Second)
}

func ExampleConfigClient_GetConfigContent() {
	client := config_client.ConfigClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	content, _ := client.GetConfigContent("TEST2", "TEST1")
	fmt.Println(content)
}
