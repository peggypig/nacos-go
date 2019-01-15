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
	Ip          string            `param:"ip"`
	Port        uint64            `param:"port"`
	Tenant      string            `param:"tenant"`
	Weight      float64           `param:"weight"`
	Enable      bool              `param:"enable"`
	Healthy     bool              `param:"healthy"`
	Metadata    map[string]string `param:"metadata"`
	ClusterName string            `param:"clusterName"`
	ServiceName string            `param:"serviceName"`
}

type LogoutServiceInstanceParam struct {
	Ip          string `param:"ip"`
	Port        uint64 `param:"port"`
	Tenant      string `param:"tenant"`
	Cluster     string `param:"cluster"`
	ServiceName string `param:"serviceName"`
}

type ModifyServiceInstanceParam struct {
	ServiceName string            `param:"serviceName"`
	Ip          string            `param:"ip"`
	Port        uint64            `param:"port"`
	Cluster     string            `param:"cluster"`
	Tenant      string            `param:"tenant"`
	Weight      float64           `param:"weight"`
	Metadata    map[string]string `param:"metadata"`
}

type GetServiceParam struct {
	Tenant      string   `param:"tenant"`
	HealthyOnly bool     `param:"healthyOnly"`
	Clusters    []string `param:"clusters"`
	ServiceName string   `param:"serviceName"`
}

type GetServiceInstanceParam struct {
	Tenant      string `param:"tenant"`
	HealthyOnly bool   `param:"healthyOnly"`
	Cluster     string `param:"cluster"`
	ServiceName string `param:"serviceName"`
	Ip          string `param:"ip"`
	Port        uint64 `param:"port"`
}

type BeatTaskParam struct {
	Ip       string            `json:"ip"`
	Port     uint64            `json:"port"`
	Weight   float64           `json:"weight"`
	Dom      string            `json:"dom"` // Dom == ServiceName
	Cluster  string            `json:"cluster"`
	Metadata map[string]string `json:"metadata"`
}

type GetServiceDetailParam struct {
	ServiceName string `param:"serviceName"`
}
