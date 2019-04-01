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
	KEY_ENDPOINT           = "endpoint"
	KEY_NAME_SPACE         = "namespace"
	KEY_ACCESS_KEY         = "accessKey"
	KEY_SECRET_KEY         = "secretKey"
	KEY_SERVER_ADDR        = "serverAddr"
	KEY_CONTEXT_PATH       = "contextPath"
	KEY_ENCODE             = "encode"
	KEY_DATA_ID            = "dataId"
	KEY_GROUP              = "group"
	KEY_TENANT             = "tenant"
	KEY_DESC               = "desc"
	KEY_APP_NAME           = "appName"
	KEY_CONTENT            = "content"
	KEY_TIMEOUT_MS         = "timeoutMs"
	KEY_LISTEN_INTERVAL    = "listenInterval"
	KEY_SERVER_CONFIGS     = "serverConfigs"
	KEY_CLIENT_CONFIG      = "clientConfig"
	CONFIG_BASE_PATH       = "/v1/cs"
	CONFIG_PATH            = CONFIG_BASE_PATH + "/configs"
	CONFIG_LISTEN_PATH     = CONFIG_BASE_PATH + "/configs/listener"
	SERVICE_BASE_PATH      = "/v1/ns"
	SERVICE_PATH           = SERVICE_BASE_PATH + "/instance"
	SERVICE_SUBSCRIBE_PATH = SERVICE_PATH + "/list"
	SERVICE_CATALOG_PATH   = SERVICE_BASE_PATH + "/catalog"
	NAMESPACE_PATH         = "/v1/console/namespaces"
	SPLIT_CONFIG           = string(rune(1))
	SPLIT_CONFIG_INNER     = string(rune(2))
	KEY_LISTEN_CONFIGS     = "Listening-Configs"
	KEY_SERVICE_NAME       = "serviceName"
	KEY_IP                 = "ip"
	KEY_PORT               = "port"
	KEY_WEIGHT             = "weight"
	KEY_ENABLE             = "enable"
	KEY_HEALTHY            = "healthy"
	KEY_METADATA           = "metadata"
	KEY_CLUSTER_NAME       = "clusterName"
	KEY_CLUSTER            = "cluster"
	KEY_BEAT               = "beat"
	KEY_DOM                = "dom"
	DEFAULT_CONTEXT_PATH   = "/nacos"
)
