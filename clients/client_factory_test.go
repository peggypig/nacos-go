package clients

import (
	"github.com/peggypig/nacos-go/clients/config_client"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/clients/service_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/stretchr/testify/assert"
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

func TestCreateConfigClientWithoutConfig(t *testing.T) {
	_, err := CreateConfigClient(map[string]interface{}{

	})
	assert.NotNil(t,err)
}

func TestCreateConfigClient(t *testing.T) {
	cfg := constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}
	client, err := CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			cfg,
		},
	})
	assert.Nil(t,err)
	realClient := client.(*config_client.ConfigClient)
	configs, _ := realClient.GetServerConfig()
	cfg.ContextPath = "/nacos"
	assert.Equal(t, 1, len(configs))
	assert.Equal(t, cfg, configs[0])
}

func TestCreateServiceClientWithoutConfig(t *testing.T) {
	_, err := CreateServiceClient(map[string]interface{}{

	})
	assert.NotNil(t,err)
}

func TestCreateServiceClient(t *testing.T) {
	cfg := constant.ServerConfig{
		IpAddr: "10.0.0.8",
		Port:   8848,
	}
	client, err := CreateServiceClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			cfg,
		},
	})
	assert.Nil(t,err)
	realClient := client.(*service_client.ServiceClient)
	configs, _ := realClient.GetServerConfig()
	cfg.ContextPath = "/nacos"
	assert.Equal(t, 1, len(configs))
	assert.Equal(t, cfg, configs[0])
}

func TestSetConfig(t *testing.T) {
	clientConfig := constant.ClientConfig{
		TimeoutMs: 10 * 1000,
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "console.nacos.io",
			Port:   80,
		},
	}
	client, err := setConfig(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	assert.Nil(t,err)
	clientConfig.BeatInterval = 5 * 1000
	clientConfig.ListenInterval = 10 * 1000
	cc, _ := client.(*nacos_client.NacosClient).GetClientConfig()
	sc, _ := client.(*nacos_client.NacosClient).GetServerConfig()
	assert.Equal(t, clientConfig, cc)
	assert.Equal(t,1,len(serverConfigs))
	assert.Equal(t, serverConfigs[0], sc[0])
}
