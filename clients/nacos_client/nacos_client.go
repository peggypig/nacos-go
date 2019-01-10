package nacos_client

import (
	"errors"
	"nacos-go/common/constant"
	"nacos-go/common/http_agent"
	"strconv"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 16:39
**/

type NacosClient struct {
	clientConfigValid  bool
	serverConfigsValid bool
	agent              http_agent.IHttpAgent
	clientConfig       constant.ClientConfig
	serverConfigs      []constant.ServerConfig
}

// 设置 clientConfig
func (client *NacosClient) SetClientConfig(config constant.ClientConfig) (err error) {
	if config.TimeoutMs <= 0 {
		err = errors.New("[client.SetClientConfig] config.TimeoutMs should > 0")
	}
	if err == nil && config.TimeoutMs <= config.ListenInterval {
		err = errors.New("[client.SetClientConfig] config.TimeoutMs should > config.ListenInterval")
	}
	if err == nil {
		if config.BeatInterval <= 0 {
			config.BeatInterval = 5 * 1000
		}
		if config.ListenInterval < 10*1000 {
			config.ListenInterval = 10 * 1000
		}
	}
	if err == nil {
		client.clientConfig = config
		client.clientConfigValid = true
	}
	return
}

// 设置 serverConfigs
func (client *NacosClient) SetServerConfig(configs []constant.ServerConfig) (err error) {
	if len(configs) <= 0 {
		err = errors.New("[client.SetServerConfig] configs can not be empty")
	}
	if err == nil {
		for index, config := range configs {
			if len(config.IpAddr) <= 0 || config.Port <= 0 || config.Port > 65535 {
				err = errors.New("[client.SetServerConfig] configs[" + strconv.Itoa(index) + "] is invalid")
				break
			}
		}
	}
	if err == nil {
		client.serverConfigs = configs
		client.serverConfigsValid = true
	}
	return
}

// 获取 clientConfig
func (client *NacosClient) GetClientConfig() (config constant.ClientConfig, err error) {
	config = client.clientConfig
	if !client.clientConfigValid {
		err = errors.New("[client.GetClientConfig] invalid client config")
	}
	return
}

// 获取serverConfigs
func (client *NacosClient) GetServerConfig() (configs []constant.ServerConfig, err error) {
	configs = client.serverConfigs
	if !client.serverConfigsValid {
		err = errors.New("[client.GetServerConfig] invalid server configs")
	}
	return
}

func (client *NacosClient) SetHttpAgent(agent http_agent.IHttpAgent) (err error) {
	if agent == nil {
		err = errors.New("[client.SetHttpAgent] http agent can not be nil")
	} else {
		client.agent = agent
	}
	return
}

func (client *NacosClient) GetHttpAgent() (agent http_agent.IHttpAgent, err error) {
	if client.agent == nil {
		err = errors.New("[client.GetHttpAgent] invalid http agent")
	} else {
		agent = client.agent
	}
	return
}
