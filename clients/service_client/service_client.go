package service_client

import "nacos-go/vo"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

type ServiceClient struct {
}

// 注册服务实例
func (client *ServiceClient) RegisterService(param vo.RegisterServiceInstanceParam) (success bool, err error) {
	return
}

// 注销服务实例
func (client *ServiceClient) LogoutService(param vo.LogoutServiceInstanceParam) (success bool, err error) {
	return
}

// 修改服务实例
func (client *ServiceClient) ModifyService(param vo.ModifyServiceInstanceParam) (success bool, err error) {
	return
}

// 获取服务列表
func (client *ServiceClient) GetService(param vo.GetServiceParam) (service vo.Service, err error) {
	return
}

// 获取服务某个实例
func (client *ServiceClient) GetServiceInstance(param vo.GetServiceInstanceParam) (serviceInstance vo.ServiceInstance, err error) {
	return
}
