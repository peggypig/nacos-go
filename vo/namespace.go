package vo

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-14 19:19
**/

type Namespace struct {
	ConfigCount       uint64 `json:"configCount"`
	Namespace         string `json:"namespace"` // namespaceId
	NamespaceShowName string `json:"namespaceShowName"`
	Quota             uint64 `json:"quota"`
	Type              uint64 `json:"type"`
}

type GetNamespaceResponse struct {
	Code    uint64      `json:"code"`
	Message string      `json:"message"`
	Data    []Namespace `json:"data"`
}
