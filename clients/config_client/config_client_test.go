package config_client

import (
	"nacos-go/common/constant"
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
	content, err := client.GetConfig("TEST", "TEST")
	if err != nil {
		t.Error(err)
	}else {
		t.Log(content)
	}
}