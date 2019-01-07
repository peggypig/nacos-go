package vo

import (
	"fmt"
	"nacos-go/common/constant"
	"nacos-go/common/util"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 18:53
**/

type ServerListManager struct {
	/**
	* 不同环境的名称
	 */
	Name                     string
	Namespace                string
	Tenant                   string
	isFixed                  bool
	isStarted                bool
	Endpoint                 string
	EndpointPort             int
	ContentPath              string
	ServerListName           string
	ServerUrls               []string
	CurrentServerAddr        string
	ServerPort               string
	AddressServerUrl         string
	initServerListRetryTimes int
	timeout                  int
}

// DEFAULT_NAME = "default"
// CUSTOM_NAME = "custom"
// FIXED_NAME = "fixed"
//和其他server的连接超时和socket超时
// initServerListRetryTimes = 5
// TIMEOUT = 5000
// endpointPort = 8080

func (manager *ServerListManager) Init(properties map[string]string) {
	manager.EndpointPort = 8080
	manager.timeout = 5000
	manager.initServerListRetryTimes = 5
	manager.isStarted = false
	serverAddrsStr := properties[constant.KEY_SERVER_ADDR]
	namespace := properties[constant.KEY_NAME_SPACE]
	endpointTmp := properties[constant.KEY_ENDPOINT]
	if len(endpointTmp) > 0 {
		manager.Endpoint = endpointTmp
	}
	contentPathTmp := properties[constant.KEY_CONTEXT_PATH]
	if len(contentPathTmp) > 0 {
		manager.ContentPath = contentPathTmp
	}
	serverListNameTmp := properties[constant.KEY_CLUSTER_NAME]
	if len(serverListNameTmp) > 0 {
		manager.ServerListName = serverListNameTmp
	}

	if len(serverAddrsStr) > 0 {
		manager.isFixed = true
		var serverAddrs []string
		var serverAddrsArr = strings.Split(serverAddrsStr, ",")
		for _, serverAddr := range serverAddrsArr {
			serverAddrArr := strings.Split(serverAddr, ":")
			if len(serverAddrArr) == 1 {
				serverAddrs = append(serverAddrs, serverAddrArr[0]+":"+util.DefaultServerPort)
			} else {
				serverAddrs = append(serverAddrs, serverAddr)
			}
		}
		manager.ServerUrls = serverAddrs
		if len(namespace) <= 0 {
			manager.Name = "fixed-" + getFixedNameSuffix(manager.ServerUrls...)
		} else {
			manager.Namespace = namespace
			manager.Tenant = namespace
			manager.Name = "fixed-" + getFixedNameSuffix(manager.ServerUrls...) + "-" + namespace
		}
	} else {
		if len(manager.Endpoint) <= 0 {
			panic("endpoint is blank")
		}
		manager.isFixed = false
		if len(namespace) <= 0 {
			manager.Name = manager.Endpoint
			manager.AddressServerUrl = fmt.Sprintf("http://%s:%d/%s/%s", manager.Endpoint, manager.EndpointPort, manager.ContentPath,
				manager.ServerListName)
		} else {
			manager.Namespace = namespace
			manager.Tenant = namespace
			manager.Name = manager.Endpoint + "-" + namespace
			manager.AddressServerUrl = fmt.Sprintf("http://%s:%d/%s/%s?namespace=%s", manager.Endpoint, manager.EndpointPort,
				manager.ContentPath, manager.ServerListName, namespace)
		}
	}
}

func getFixedNameSuffix(serverIps ...string) (sb string) {
	var split = ""
	for _, serverIp := range serverIps {
		sb += split
		sb += strings.Replace(serverIp, ":", "_", -1)
		split = "-"
	}
	return
}
