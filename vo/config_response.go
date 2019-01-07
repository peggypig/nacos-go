package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 16:32
**/

type ConfigResponse struct {
	param   map[string]interface{}
	context IConfigContext
}

func (response *ConfigResponse) checkParam() {
	if response.param == nil {
		response.param = make(map[string]interface{})
	}
}

func (response *ConfigResponse) GetConfigContext() (context IConfigContext) {
	return response.context
}

func (response *ConfigResponse) GetParameter(key string) (value interface{}) {
	response.checkParam()
	return response.param[key]
}

func (response *ConfigResponse) GetTenant() (tenant string) {
	response.checkParam()
	return response.param["tenant"].(string)
}

func (response *ConfigResponse) SetTenant(tenant string) {
	response.checkParam()
	response.param["tenant"] = tenant
}

func (response *ConfigResponse) SetDataId(dataId string) {
	response.checkParam()
	response.param["dataId"] = dataId
}

func (response *ConfigResponse) GetDataId() (dataId string) {
	response.checkParam()
	return response.param["dataId"].(string)
}

func (response *ConfigResponse) SetGroup(group string) {
	response.checkParam()
	response.param["group"] = group
}

func (response *ConfigResponse) GetGroup() (group string) {
	response.checkParam()
	return response.param["group"].(string)
}

func (response *ConfigResponse) SetContent(content string) {
	response.checkParam()
	response.param["content"] = content
}

func (response *ConfigResponse) GetContent() (group string) {
	response.checkParam()
	return response.param["content"].(string)
}
