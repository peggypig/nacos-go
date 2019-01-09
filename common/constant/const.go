package constant

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

const (
	KEY_ENDPOINT        = "endpoint"
	KEY_NAME_SPACE      = "namespace"
	KEY_ACCESS_KEY      = "accessKey"
	KEY_SECRET_KEY      = "secretKey"
	KEY_SERVER_ADDR     = "serverAddr"
	KEY_CONTEXT_PATH    = "contextPath"
	KEY_CLUSTER_NAME    = "clusterName"
	KEY_ENCODE          = "encode"
	KEY_DATA_ID         = "dataId"
	KEY_GROUP           = "group"
	KEY_TENANT          = "tenant"
	KEY_DESC            = "desc"
	KEY_APP_NAME        = "appName"
	KEY_CONTENT         = "content"
	KEY_TIMEOUT_MS      = "timeoutMs"
	KEY_LISTEN_INTERVAL = "listenInterval"
	KEY_SERVER_CONFIG   = "serverConfig"
	CONFIG_BASE_PATH    = "/nacos/v1/cs"
	CONFIG_PATH         = CONFIG_BASE_PATH + "/configs"
	CONFIG_LISTEN_PATH  = CONFIG_BASE_PATH + "/configs/listener"
	SERVICE_BASE_PATH   = "/nacos/v1/ns/"
	SERVICE_PATH        = SERVICE_BASE_PATH + "instance"
	SPLIT_CONFIG        = string(rune(1))
	SPLIT_CONFIG_INNER  = string(rune(2))
	KEY_LISTEN_CONFIGS  = "Listening-Configs"
)
