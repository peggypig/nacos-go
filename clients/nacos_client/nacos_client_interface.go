package nacos_client

import (
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 16:32
**/

//go:generate mockgen -destination ../../mock/mock_nacos_client_interface.go -package mock -source=./nacos_client_interface.go

type INacosClient interface {
	SetClientConfig(constant.ClientConfig) error
	SetServerConfig([]constant.ServerConfig) error
	GetClientConfig() (constant.ClientConfig, error)
	GetServerConfig() ([]constant.ServerConfig, error)
	SetHttpAgent(http_agent.IHttpAgent) error
	GetHttpAgent() (http_agent.IHttpAgent, error)

	// namespace
	GetNamespace() ([]vo.Namespace, error)
	CreateNamespace(vo.CreateNamespaceParam) (bool, error)
	ModifyNamespace(vo.ModifyNamespaceParam) (bool, error)
	DeleteNamespace(vo.DeleteNamespaceParam) (bool, error)
}
