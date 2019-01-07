package vo

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-07 16:34
**/

type IConfigContext interface {
	// get param
	GetParameter(key string) interface{}

	// set context
	SetParameter(key string,value interface{})
}