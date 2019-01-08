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
	// 获取配置
	// dataId  require
	// group   require
	// tenant ==>nacos.namespace optional
	GetConfig(param vo.ConfigParam) (string, error)

	// 发布配置
	// dataId  require
	// group   require
	// content require
	// tenant ==>nacos.namespace optional
	PublishConfig(param vo.ConfigParam) (bool, error)

	// 删除配置
	// dataId  require
	// group   require
	// tenant ==>nacos.namespace optional
	DeleteConfig(param vo.ConfigParam) (bool, error)

	// 监听配置
	// dataId  require
	// group   require
	// tenant ==>nacos.namespace optional
	ListenConfig(params []vo.ConfigParam) (err error)

	// 停止监听配置变化
	StopListenConfig()

	// 获取配置内容
	// dataId  require
	// group   require
	// 先从本地监听的配置中获取，没有则从服务器上获取
	GetConfigContent(dataId string,groupId string)(string,error)
}
