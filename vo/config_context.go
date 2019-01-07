package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 16:38
**/

type ConfigContext struct {
	param map[string]interface{}
}

func (context *ConfigContext) SetParameter(key string, value interface{}) {
	if context.param == nil {
		context.param = make(map[string]interface{})
	}
	context.param[key] = value
}

func (context *ConfigContext) GetParameter(key string) (value interface{}) {
	if context.param == nil {
		context.param = make(map[string]interface{})
	}
	return context.param[key]
}
