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
	Metadata    map[string]string
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
	Metadata    map[string]string
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

type BeatTaskParam struct {
	Ip       string            `json:"ip"`
	Port     uint64            `json:"port"`
	Weight   float64           `json:"weight"`
	Dom      string            `json:"dom"` // Dom == ServiceName
	Cluster  string            `json:"cluster"`
	MetaData map[string]string `json:"metaData"`
}

type GetServiceDetailParam struct {
	ServiceName string
}
