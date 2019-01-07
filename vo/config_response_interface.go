package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 16:34
**/

type IConfigResponse interface {
	// get param
	GetParameter(key string) interface{}

	// get context
	GetConfigContext() IConfigContext
}
