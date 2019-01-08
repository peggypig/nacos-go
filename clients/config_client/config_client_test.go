package config_client

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
* @create : 2019-01-08 12:05
**/

func TestConfigClient_GetConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs:[]constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		},},
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId:"TEST",
		Group:"DEFAULT_GROUP",
	})
	if err != nil {
		t.Error(err)
	}else {
		t.Log(content)
	}
}


func TestConfigClient_PublishConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs:[]constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		},},
	}
	content, err := client.PublishConfig(vo.ConfigParam{
		DataId:"TEST",
		Group:"DEFAULT_GROUP",
		Content:"aaa",
		Tenant:"bbb",
	})
	if err != nil {
		t.Error(err)
	}else {
		t.Log(content)
	}
}