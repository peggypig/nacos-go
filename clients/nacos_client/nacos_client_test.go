package nacos_client

import (
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-15 11:59
**/

func TestNacosClient_SetClientConfig(t *testing.T) {
	client := NacosClient{}
	cfg := constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	}
	err := client.SetClientConfig(cfg)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, cfg, client.clientConfig)
}

func TestNacosClient_SetServerConfig(t *testing.T) {
	client := NacosClient{}
	cfg := constant.ServerConfig{
		ContextPath: "/nacos",
		IpAddr:      "console.nacos.io",
		Port:        8848,
	}
	err := client.SetServerConfig([]constant.ServerConfig{cfg})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(client.serverConfigs))
	assert.Equal(t, cfg,client.serverConfigs[0])
}



func TestNacosClient_SetHttpAgent(t *testing.T) {
	client := NacosClient{}
	agent := http_agent.HttpAgent{}
	err := client.SetHttpAgent(&agent)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, &agent, client.agent)
}
