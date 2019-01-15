package clients

import (
	"errors"
	"github.com/peggypig/nacos-go/clients/config_client"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/clients/service_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
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
func CreateConfigClient(properties map[string]interface{}) (iClient config_client.IConfigClient,
	err error) {
	client := config_client.ConfigClient{}
	nacosClient, errSetConfig := setConfig(properties)
	if errSetConfig != nil {
		err = errSetConfig
	}
	if err == nil {
		client.INacosClient = nacosClient
		_ = client.SetHttpAgent(&http_agent.HttpAgent{})
		iClient = &client
	}
	return
}

// 创建服务发现相关的客户端
func CreateServiceClient(properties map[string]interface{}) (iClient service_client.IServiceClient,
	err error) {
	client := service_client.ServiceClient{}
	nacosClient, errSetConfig := setConfig(properties)
	if errSetConfig != nil {
		err = errSetConfig
	}
	if err == nil {
		client.INacosClient = nacosClient
		_ = client.SetHttpAgent(&http_agent.HttpAgent{})
		iClient = &client
	}
	return
}

func setConfig(properties map[string]interface{}) (iClient nacos_client.INacosClient,
	err error) {
	client := nacos_client.NacosClient{}
	if clientConfigTmp, exist := properties[constant.KEY_CLIENT_CONFIG]; exist {
		if clientConfig, ok := clientConfigTmp.(constant.ClientConfig); ok {
			err = client.SetClientConfig(clientConfig)
		}
	} else {
		_ = client.SetClientConfig(constant.ClientConfig{
			TimeoutMs:      30 * 1000,
			ListenInterval: 10 * 1000,
			BeatInterval:   5 * 1000,
		})
	}
	// 设置 serverConfig
	if err == nil {
		if serverConfigTmp, exist := properties[constant.KEY_SERVER_CONFIGS]; exist {
			if serverConfigs, ok := serverConfigTmp.([]constant.ServerConfig); ok {
				err = client.SetServerConfig(serverConfigs)
			}
		} else {
			err = errors.New("server configs not found in properties")
		}
	}
	if err == nil {
		iClient = &client
	}
	return
}
