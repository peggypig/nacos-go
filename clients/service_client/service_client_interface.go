package service_client

import (
	"github.com/peggypig/nacos-go/vo"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 09:56
**/

//go:generate mockgen -destination ../../mock/mock_service_client_interface.go -package mock -source=./service_client_interface.go

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
	// 开始发送心跳的任务  只有在service.healthCheckMode = client的情况下才有效
	StartBeatTask(param vo.BeatTaskParam) error
	// 停止发送心跳的任务  当前一次发送心跳任务完成后  结束
	StopBeatTask()
	// 获取service的基本信息
	GetServiceDetail(param vo.GetServiceDetailParam) (vo.ServiceDetail, error)
	// 服务监听
	Subscribe(param vo.SubscribeParam) error
	// 查询全局服务列表
	GetServiceList(param vo.GetServiceListParam) (vo.ServiceSummaryList, error)
}
