package example

import (
	"fmt"
	"github.com/peggypig/nacos-go/clients"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/vo"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 11:25
**/


func ExampleCreateConfigClient_GetConfig() {
	client, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "console.nacos.io",
				Port:   80,
			},
		},
	})
	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	fmt.Println(content)
}

func ExampleCreateServiceClient_GetServiceDetail() {
	client, _ := clients.CreateServiceClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "console.nacos.io",
				Port:   80,
			},
		},
	})
	service, _ := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "demoservice",
	})
	fmt.Println(service)
}

