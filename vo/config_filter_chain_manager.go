package vo

import "sync"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 18:36
**/

type ConfigFilterChainManager struct {
	filters []IConfigFilter
	mutex   sync.Mutex
}

func (manager *ConfigFilterChainManager) ConfigFilterChainManager(filter IConfigFilter) *ConfigFilterChainManager {
	// 根据order大小顺序插入
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	i := 0
	for i < len(manager.filters) {
		var currentValue = manager.filters[i]
		if currentValue.GetFilterName() == filter.GetFilterName() {
			break
		}
		if filter.GetOrder() >= currentValue.GetOrder() && i < len(manager.filters) {
			i++
		} else {
			prefix := manager.filters[0:i]
			suffix := manager.filters[i:]
			manager.filters = append(append(prefix, currentValue), suffix...)
			break
		}
	}
	if i == len(manager.filters) {
		prefix := manager.filters[0:i]
		suffix := manager.filters[i:]
		manager.filters = append(append(prefix, filter), suffix...)
	}
	return manager
}

func (manager *ConfigFilterChainManager) DoFilter(request IConfigRequest, response IConfigResponse) {
	virtualConfigFilterChain{}.init(manager.filters).DoFilter(request, response)
}

type virtualConfigFilterChain struct {
	additionalFilters []IConfigFilter
	currentPosition   int
}

func (chain *virtualConfigFilterChain) init(filters []IConfigFilter) *virtualConfigFilterChain {
	chain.additionalFilters = filters
	return chain
}

func (chain *virtualConfigFilterChain) DoFilter(request IConfigRequest, response IConfigResponse) {
	if chain.currentPosition != len(chain.additionalFilters) {
		chain.currentPosition++
		var nextFilter = chain.additionalFilters[chain.currentPosition-1]
		nextFilter.DoFilter(request, response, chain)
	}
}
