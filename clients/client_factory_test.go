package clients

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
* @create : 2019-01-08 11:25
**/

func TestCreateConfigClient(t *testing.T) {
	client, err := CreateConfigClient(map[string]interface{}{
		"serverConfig": constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		},
	})
	if err != nil {
		t.Error("error:", err)
	} else {
		t.Log("client:", client)
	}
}