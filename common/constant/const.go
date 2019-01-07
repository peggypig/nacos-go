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
	KEY_ENDPOINT                  = "endpoint"
	KEY_NAME_SPACE                 = "namespace"
	KEY_ACCESSS_KEY                = "accessKey"
	KEY_SECRET_KEY                 = "secretKey"
	KEY_SERVER_ADDR                = "serverAddr"
	KEY_CONTEXT_PATH               = "contextPath"
	KEY_CLUSTER_NAME               = "clusterName"
	KEY_ENCODE                     = "encode"
	KEY_NAMING_LOAD_CACHE_AT_START = "namingLoadCacheAtStart"

	CLIENT_VERSION_HEADER = "Client-Version"

	CLIENT_VERSION = "3.0.0"

	BODY_VERSION = 204

	DEFAULT_GROUP = "DEFAULT_GROUP"

	APPNAME = "AppName"

	UNKNOWN_APP = "UnknownApp"

	DEFAULT_DOMAINNAME = "commonconfig.config-host.taobao.com"

	DAILY_DOMAINNAME = "commonconfig.taobao.net"

	NULL = ""

	DATAID = "dataId"

	GROUP = "group"

	LAST_MODIFIED = "Last-Modified"

	ACCEPT_ENCODING = "Accept-Encoding"

	CONTENT_ENCODING = "Content-Encoding"

	PROBE_MODIFY_REQUEST = "Listening-Configs"

	PROBE_MODIFY_RESPONSE = "Probe-Modify-Response"

	PROBE_MODIFY_RESPONSE_NEW = "Probe-Modify-Response-New"

	USE_ZIP = "true"

	CONTENT_MD5 = "Content-MD5"

	CONFIG_VERSION    = "Config-Version"
	IF_MODIFIED_SINCE = "If-Modified-Since"

	SPACING_INTERVAL = "client-spacing-interval"

	BASE_PATH = "/v1/cs"

	CONFIG_CONTROLLER_PATH = BASE_PATH + "/configs"
	/**
	 * second
	 */
	ASYNC_UPDATE_ADDRESS_INTERVAL = 300

	/**
	 * second
	 */
	POLLING_INTERVAL_TIME = 15

	/**
	 * millisecond
	 */
	ONCE_TIMEOUT = 2000

	/**
	 * millisecond
	 */
	CONN_TIMEOUT = 2000

	/**
	 * millisecond
	 */
	SO_TIMEOUT = 60000

	/**
	 * millisecond
	 */
	RECV_WAIT_TIMEOUT = ONCE_TIMEOUT * 5

	ENCODE = "UTF-8"

	MAP_FILE = "map-file.js"

	FLOW_CONTROL_THRESHOLD = 20

	FLOW_CONTROL_SLOT = 10

	FLOW_CONTROL_INTERVAL = 1000

	LINE_SEPARATOR = '1'

	WORD_SEPARATOR = '2'

	LONGPULLING_LINE_SEPARATOR = "\r\n"

	CLIENT_APPNAME_HEADER       = "Client-AppName"
	CLIENT_REQUEST_TS_HEADER    = "Client-RequestTS"
	CLIENT_REQUEST_TOKEN_HEADER = "Client-RequestToken"

	ATOMIC_MAX_SIZE = 1000

	NAMING_INSTANCE_ID_SPLITTER  = "#"
	NAMING_INSTANCE_ID_SEG_COUNT = 4
	NAMING_HTTP_HEADER_SPILIER   = "\\|"

	NAMING_DEFAULT_CLUSTER_NAME = "DEFAULT"
)
