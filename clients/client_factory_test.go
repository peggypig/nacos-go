package clients

import (
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/vo"
	"testing"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 11:25
**/

func TestCreateConfigClient(t *testing.T) {
	client, err := CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "10.0.0.8",
				Port:   8848,
			},
		},
	})
	if err != nil {
		t.Error("error:", err)
	} else {
		t.Log("client:", client)
	}
}

func TestCreateConfigClient_GetConfig(t *testing.T) {
	client, err := CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "10.0.0.8",
				Port:   8848,
			},
		},
	})
	if err != nil {
		t.Error("error:", err)
	} else {
		t.Log("client:", client)
	}
	if err == nil {
		content, errGet := client.GetConfig(vo.ConfigParam{
			DataId: "TEST",
			Group:  "TEST",
		})
		if errGet != nil {
			t.Error(errGet)
		} else {
			t.Log(content)
		}
	}
}

func TestCreateServiceClient_GetServiceDetail(t *testing.T) {
	client, err := CreateServiceClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "10.0.0.8",
				Port:   8848,
			},
		},
	})
	if err != nil {
		t.Error("error:", err)
	} else {
		t.Log("client:", client)
	}
	if err == nil {
		service, errGet := client.GetServiceDetail(vo.GetServiceDetailParam{
			ServiceName: "demoservice",
		})
		if errGet != nil {
			t.Error(errGet)
		} else {
			t.Logf("%+v", service)
		}
	}
}

func TestSetConfig(t *testing.T) {
	client, err := setConfig(map[string]interface{}{
		"clientConfig": constant.ClientConfig{
			TimeoutMs: 10 * 1000,
		},
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "console.nacos.io",
				Port:   80,
			},
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(client.GetServerConfig())
		t.Log(client.GetClientConfig())
	}
}
