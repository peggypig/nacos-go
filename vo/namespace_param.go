package vo

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-14 19:23
**/

type CreateNamespaceParam struct {
	NamespaceName string `param:"namespaceName"`
	NamespaceDesc string `param:"namespaceDesc"`
}

type ModifyNamespaceParam struct {
	Namespace     string `param:"namespace"` // NamespaceId
	NamespaceName string `param:"namespaceShowName"`
	NamespaceDesc string `param:"namespaceDesc"`
}

type DeleteNamespaceParam struct {
	NamespaceId string `param:"namespaceId"`
}