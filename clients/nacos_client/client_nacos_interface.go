package nacos_client

import "nacos-go/common/constant"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 16:32
**/

type INacosClient interface {
	SetClientConfig(constant.ClientConfig) error
	SetServerConfig([]constant.ServerConfig) error
	GetClientConfig() (constant.ClientConfig, error)
	GetServerConfig() ([]constant.ServerConfig, error)
}
