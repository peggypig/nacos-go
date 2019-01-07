package util

import (
	"log"
	"os"
	"strconv"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 18:58
**/

var DefaultContextPath = "nacos"
var DefaultNodesPath = "serverlist"
var AppKey string
var AppName string
var DefaultServerPort string
var ClientVersion string
var ConnectTimeout int
var PerTaskConfigSize int

func init() {
	// 客户端身份信息
	AppKey = os.Getenv("nacos.client.appKey")
	AppName = os.Getenv("project.name")
	DefaultServerPort = os.Getenv("nacos.server.port")
	if len(DefaultServerPort) <= 0 {
		DefaultServerPort = "8848"
	}
	log.Println("settings", "nacos-server port:", defaultServerPort)
	connectTimeoutTemp := os.Getenv("NACOS.CONNECT.TIMEOUT")
	if len(connectTimeoutTemp) <= 0 {
		ConnectTimeout = 1000
	} else {
		i, err := strconv.Atoi(connectTimeoutTemp)
		if err != nil {
			log.Println("settings", "invalid NACOS.CONNECT.TIMEOUT:", connectTimeoutTemp)
		} else {
			ConnectTimeout = i
		}
	}
	ClientVersion = os.Getenv("nacos.client.version")
	if len(ClientVersion) <= 0 {
		ClientVersion = "unknown"
	}
	perTaskConfigSizeTemp := os.Getenv("PER_TASK_CONFIG_SIZE")
	if len(perTaskConfigSizeTemp) <= 0 {
		PerTaskConfigSize = 3000
	} else {
		i, err := strconv.Atoi(perTaskConfigSizeTemp)
		if err != nil {
			log.Println("settings", "invalid env PER_TASK_CONFIG_SIZE:", perTaskConfigSizeTemp)
		} else {
			PerTaskConfigSize = i
		}
	}
}
