package service_client

import (
	"nacos-go/common/constant"
	"nacos-go/vo"
	"testing"
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
	client := ServiceClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs: 30 * 1000,
		},
	}
	success, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demoservice",
		Weight:      -1,
		ClusterName: "a",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(success)
	}
}

func TestServiceClient_ModifyServiceInstance(t *testing.T) {
	client := ServiceClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs: 30 * 1000,
		},
	}
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
	client := ServiceClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs: 30 * 1000,
		},
	}
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
	client := ServiceClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs: 30 * 1000,
		},
	}
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
	client := ServiceClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs: 30 * 1000,
		},
	}
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

