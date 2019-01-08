package config_client

import "nacos-go/vo"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 10:03
**/

type IConfigClient interface {
	GetConfig(param vo.ConfigParam) (string, error)
	PublishConfig(param vo.ConfigParam) (bool, error)
	DeleteConfig(param vo.ConfigParam) (bool, error)
	ListenConfig()
}
