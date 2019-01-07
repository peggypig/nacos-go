package config_client

import (
	"nacos-go/common/constant"
	"nacos-go/common/http"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:40
**/

type ClientWorker struct {
	Agent http.HttpAgent
}

func (worker *ClientWorker) AddListeners() {

}

func (worker *ClientWorker) RemoveListener() {

}

func (worker *ClientWorker) AddTenantListeners() {

}

func (worker *ClientWorker) RemoveTenantListeners() {

}

func (worker *ClientWorker) RemoveCache() {

}

func (worker *ClientWorker) RemoveCacheTenant() {

}

func (worker *ClientWorker) AddCacheDataIfAbsent() {

}

func (worker *ClientWorker) AddCacheDataIfAbsentTenant() {

}

func (worker *ClientWorker) GetCache() {

}

func (worker *ClientWorker) GetCacheTenant() {

}

func (worker *ClientWorker) GetServerConfig() {

}

func (worker *ClientWorker) checkLocalConfig() {

}

func (worker *ClientWorker) Nil2DefaultGroup() (group string) {
	return constant.DEFAULT_GROUP
}

func (worker *ClientWorker) CheckConfigInfo() {

}

// 从Server获取值变化了的DataID列表。返回的对象里只有dataId和group是有效的。
// 保证不返回nil
func (worker *ClientWorker) CheckUpdateDataIds() {

}

// 从Server获取值变化了的DataID列表。返回的对象里只有dataId和group是有效的。
// 保证不返回nil
func (worker *ClientWorker) CheckUpdateConfigStr() {

}

// 从HTTP响应拿到变化的groupKey。保证不返回NULL。
func (worker *ClientWorker) ParseUpdateDataIdResponse() {

}

func (worker *ClientWorker) IsHealthServer() {

}
