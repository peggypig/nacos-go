package vo

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-07 18:34
**/

type IConfigFilterChain interface {
	DoFilter(request IConfigRequest,response IConfigResponse)
}