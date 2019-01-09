package vo

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 10:26
**/

type ServiceInstance struct {
	InstanceId  string      `json:"instanceId"`
	Ip          string      `json:"ip"`
	Port        uint64      `json:"port"`
	Metadata    interface{} `json:"metadata"`
	Service     string      `json:"service"`
	Healthy     bool        `json:"healthy"`
	ClusterName string      `json:"clusterName"`
	Weight      float64     `json:"weight"`
}

type Host struct {
	Valid       bool        `json:"valid"`
	Marked      bool        `json:"marked"`
	InstanceId  string      `json:"instanceId"`
	Port        uint64      `json:"port"`
	Ip          string      `json:"ip"`
	Weight      float64     `json:"weight"`
	Metadata    interface{} `json:"metadata"`
	ClusterName string      `json:"clusterName"`
	ServiceName string      `json:"serviceName"`
	Enable      bool        `json:"enable"`
}

type Service struct {
	Dom             string      `json:"dom"`
	CacheMillis     uint64      `json:"cacheMillis"`
	UseSpecifiedURL bool        `json:"useSpecifiedUrl"`
	Hosts           []Host      `json:"hosts"`
	Checksum        string      `json:"checksum"`
	LastRefTime     uint64      `json:"lastRefTime"`
	Env             string      `json:"env"`
	Clusters        string      `json:"clusters"`
	Metadata        interface{} `json:"metadata"`
}
