### nacos-go
go语言版本的nacos client，支持config_client和service_client

#### client的config
- ClientConfig 客户端配置参数  
```go
    constant.ClientConfig{
		TimeoutMs:      30 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	}
```
TimeoutMs：http请求超时时间，单位毫秒  
ListenInterval：监听间隔时间，单位毫秒（仅在ConfigClient中有效）  
BeatInterval：心跳间隔时间，单位毫秒（仅在ServiceClient中有效）

- ServerConfig nacos服务信息配置参数
```go
    constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		ContextPath: "/nacos",
		Port:        80,
	}
```
IpAddr：nacos服务的ip地址  
Port：nacos服务端口  
ContextPath：nacos服务的上下文路径，默认是“/nacos”  
<b>注：ServerConfig支持配置多个，在请求出错时，自动切换</b>

#### nacos_client
1. GetNamespace   
获取namespace
2. CreateNamespace  
创建namespace
3. ModifyNamespace  
修改namespace
4. DeleteNamespace  
删除namespace

#### config_client
1. GetConfig    
从server端获取配置
2. GetConfigContent  
会优先从本地缓存中获取，如果没有，才会从server端获取
3. PublishConfig  
发布配置到server端
4. DeleteConfig  
删除配置
5. ListenConfig   
监听配置变化
6. StopListenConfig    
关闭配置监听
7. AddConfigToListen  
增加监听配置，在ListenConfig后才会生效

#### service_client
1. RegisterServiceInstance  
注册服务实例  
2. LogoutServiceInstance  
注销服务实例  
3. ModifyServiceInstance  
修改服务实例  
4. GetService  
获取服务列表  
5. GetServiceInstance  
获取服务实例  
6. StartBeatTask  
向服务器发送健康心跳  
7. StopBeatTask  
停止向服务器发送心跳  
8. GetServiceDetail  
获取服务的详细信息  
9. Subscribe  
服务监听

### quick start
以GetConfig为例：  
Step 1. 构造相关参数  
```go
    // 可以没有，采用默认值
    clientConfig := constant.ClientConfig{
    		TimeoutMs:      30 * 1000,
    		ListenInterval: 10 * 1000,
    		BeatInterval:   5 * 1000,
    	} 
    // 至少一个
    serverConfigs := []constant.ServerConfig{
    	{
    	    IpAddr:      "console1.nacos.io",
    	    ContextPath: "/nacos",
    	    Port:        80,
        },
        {
            IpAddr:      "console2.nacos.io",
            ContextPath: "/nacos",
            Port:        80,
        },
    }
```
Step 2. 构造客户端
```go
    // 如果参数设置不合法，将抛出error
    client, err := CreateConfigClient(map[string]interface{}{
    	"serverConfigs": serverConfigs,
    	"clientConfig":  clientConfig,
    })
```
Step 3. 目标操作
```go
        if err == nil {
        	// 从服务端获取config
    		content, errGet := client.GetConfig(vo.ConfigParam{
    			DataId: "TEST",
    			Group:  "TEST",
    		})
    		if errGet != nil {
    			fmt.Println(errGet)
    		} else {
    			fmt.Println(content)
    		}
    	}
```


