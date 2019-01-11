package config_client

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
* @create : 2019-01-08 12:05
**/

func TestConfigClient_GetConfig(t *testing.T) {
	client := ConfigClient{}
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
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
		Tenant: "bbb",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_PublishConfig(t *testing.T) {
	client := ConfigClient{}
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
	content, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "DEFAULT_GROUP",
		Content: "aaa",
		Tenant:  "73d76b32-985c-49b8-8b9d-ec64e956f3d1",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_DeleteConfig(t *testing.T) {
	client := ConfigClient{}
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
	content, err := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
		Tenant: "bbb",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_ListenConfig(t *testing.T) {
	client := ConfigClient{}
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
	err := client.ListenConfig([]vo.ConfigParam{{
		DataId:  "TESTa",
		Group:   "TEST",
		Content: "2019-01-08 09:57:34",
	}, {
		DataId: "TESTa",
		Group:  "DEFAULT_GROUP",
	}})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(100 * time.Second)
}

func TestConfigClient_GetConfigContent(t *testing.T) {
	client := ConfigClient{}
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
	content, err := client.GetConfigContent("TEST", "TEST1")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}
