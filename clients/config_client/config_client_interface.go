package config_client

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 10:03
**/

type IConfigClient interface {
	GetConfig(dataId string, group string) (string, error)
	PublishConfig(dataId string, group string, content string) (bool, error)
	DeleteConfig(dataId string, group string) (bool, error)
	ListenConfig()
}
