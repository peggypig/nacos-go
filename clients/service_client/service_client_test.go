package service_client

import (
	"nacos-go/clients/nacos_client"
	"nacos-go/common/constant"
	"nacos-go/common/http_agent"
	"nacos-go/vo"
	"testing"
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

func TestServiceClient_RegisterServiceInstance(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.81",
		Port:   8848,
	},constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	success, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demoservice1",
		Weight:      1000,
		ClusterName: "a",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(success)
	}
}

func TestServiceClient_ModifyServiceInstance(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	success, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demoservice",
		Weight:      2,
		ClusterName: "a",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(success)
	}
}

func TestServiceClient_GetService(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	service, err := client.GetService(vo.GetServiceParam{
		ServiceName: "demoservice",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", service)
	}
}

func TestServiceClient_GetServiceInstance(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	service, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "demoservice",
		Ip:          "10.0.0.10",
		Port:        8848,
		Cluster:     "a",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", service)
	}
}

func TestServiceClient_LogoutServiceInstance(t *testing.T) {
	client := ServiceClient{}
	client.INacosClient = &nacos_client.NacosClient{}
	_ = client.SetHttpAgent(&http_agent.HttpAgent{})
	_ = client.SetServerConfig([]constant.ServerConfig{constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}})
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs: 30 * 1000,
	})
	service, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "demoservice",
		Ip:          "10.0.0.10",
		Port:        8848,
		Cluster:     "a",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", service)
	}
}

func TestServiceClient_StartBeatTask(t *testing.T) {
	client := ServiceClient{}
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
	err := client.StartBeatTask(vo.BeatTaskParam{
		Ip: "10.0.0.10",
		//Port:    8848,
		//Cluster: "a",
		Dom: "demoservice",
	})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(100 * time.Second)
}

func TestServiceClient_GetServiceInfo(t *testing.T) {
	client := ServiceClient{}
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
	info, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "demoservice",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", info)
	}
}
