package service_client

import "nacos-go/vo"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 09:56
**/

type IServiceClient interface {
	// 注册服务实例
	RegisterServiceInstance(param vo.RegisterServiceInstanceParam) (bool, error)
	// 注销服务实例
	LogoutServiceInstance(param vo.LogoutServiceInstanceParam) (bool, error)
	// 修改服务实例
	ModifyServiceInstance(param vo.ModifyServiceInstanceParam) (bool, error)
	// 获取服务列表
	GetService(param vo.GetServiceParam) (vo.Service, error)
	// 获取服务某个实例
	GetServiceInstance(param vo.GetServiceInstanceParam) (vo.ServiceInstance, error)
}
