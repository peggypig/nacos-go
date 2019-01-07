package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 18:29
**/

type IConfigFilter interface {
	Init(config IFilterConfig)
	DoFilter(request IConfigRequest, response IConfigResponse, filterChain IConfigFilterChain)
	Deploy()
	GetOrder() int
	GetFilterName() string
}
