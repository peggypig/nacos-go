package clients

import (
	"nacos-go/clients/config_client"
	"nacos-go/clients/service_client"
	"nacos-go/common/constant"
	"os"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

// 创建配置相关的客户端
func CreateConfigClient(properties map[string]string) (client config_client.ConfigClient) {
	client = config_client.ConfigClient{}
	if encodeTmp, exist := properties[constant.KEY_ENCODE]; exist {
		client.Encode = strings.Trim(encodeTmp, " ")
	} else {
		client.Encode = constant.ENCODE
	}
	if namespaceTmp, exist := properties[constant.KEY_NAME_SPACE]; exist {
		client.NameSpace = strings.Trim(namespaceTmp, " ")
	} else {
		userTenant := strings.Trim(os.Getenv("tenant.id"), " ")
		if len(userTenant) <= 0 {
			userTenant = strings.Trim(os.Getenv("acm.namespace"), " ")
		}
		client.NameSpace = userTenant
	}
	client.Agent.Properties = properties
	client.Agent.Start()
	client.Worker.Agent = client.Agent
	return
}

// 创建服务发现相关的客户端
func CreateServiceClient() (client service_client.ServiceClient) {
	client = service_client.ServiceClient{}
	return
}
