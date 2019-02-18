package example

import (
	"fmt"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/clients/service_client"
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
* @create : 2019-01-09 11:12
**/

func ExampleServiceClient_RegisterServiceInstance() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	success, _ := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo",
		Weight:      1000,
		ClusterName: "a",
	})
	fmt.Println(success)
}

func ExampleServiceClient_ModifyServiceInstance() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	success, _ := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demoservice",
		Weight:      2,
		ClusterName: "a",
	})
	fmt.Println(success)
}

func ExampleServiceClient_GetService() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	service, _ := client.GetService(vo.GetServiceParam{
		ServiceName: "unit",
	})
	fmt.Println(service)
}

func ExampleServiceClient_GetServiceInstance() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	service, _ := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "demoservice",
		Ip:          "10.0.0.10",
		Port:        8848,
		Cluster:     "a",
	})
	fmt.Println(service)
}

func ExampleServiceClient_LogoutServiceInstance() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	success, _ := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "demoservice",
		Ip:          "10.0.0.10",
		Port:        8848,
		Cluster:     "a",
	})
	fmt.Println(success)
}

func ExampleServiceClient_StartBeatTask() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:    30 * 1000,
		BeatInterval: 2 * 1000,
	})
	_ = client.StartBeatTask(vo.BeatTaskParam{
		Ip: "10.0.0.10",
		//Port:    8848,
		//Cluster: "a",
		Dom: "demoservice",
	})
	time.Sleep(20 * time.Second)
}

func ExampleServiceClient_GetServiceInfo() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:    30 * 1000,
		BeatInterval: 2 * 1000,
	})
	info, _ := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "demoservice",
	})
	fmt.Println(info)
}

func ExampleServiceClient_Subscribe() {
	client := service_client.ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	_ = client.Subscribe(vo.SubscribeParam{
		ServiceName: "unit",
		//Clusters:    []string{"a"},
		SubscribeCallback: func(services []vo.SubscribeService,err error) {
			fmt.Println(err)
			fmt.Println(services)
		},
	})
}
