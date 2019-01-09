package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 10:03
**/

type RegisterServiceInstanceParam struct {
	Ip          string
	Port        uint64
	Tenant      string
	Weight      float64
	Enable      bool
	Healthy     bool
	Metadata    string
	ClusterName string
	ServiceName string
}

type LogoutServiceInstanceParam struct {
	Ip          string
	Port        uint64
	Tenant      string
	Cluster     string
	ServiceName string
}

type ModifyServiceInstanceParam struct {
	ServiceName string
	Ip          string
	Port        uint64
	Cluster     string
	Tenant      string
	Weight      float64
	Metadata    string
}

type GetServiceParam struct {
	Tenant      string
	HealthyOnly bool
	Clusters    []string
	ServiceName string
}

type GetServiceInstanceParam struct {
	Tenant      string
	HealthyOnly bool
	Cluster     string
	ServiceName string
	Ip          string
	Port        uint64
}
