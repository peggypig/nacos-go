package service_client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"nacos-go/clients/nacos_client"
	"nacos-go/common/constant"
	"nacos-go/common/http_agent"
	"nacos-go/common/util"
	"nacos-go/vo"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

type ServiceClient struct {
	nacos_client.INacosClient
	beating bool
	mutex   sync.Mutex
}

// 获取参数配置
func (client *ServiceClient) sync() (clientConfig constant.ClientConfig,
	serverConfigs []constant.ServerConfig, agent http_agent.IHttpAgent, err error) {
	clientConfig, err = client.GetClientConfig()
	if err != nil {
		log.Println(err, ";do you call client.SetClientConfig()?")
	}
	if err == nil {
		serverConfigs, err = client.GetServerConfig()
		if err != nil {
			log.Println(err, ";do you call client.SetServerConfig()?")
		}
	}
	if err == nil {
		agent, err = client.GetHttpAgent()
		if err != nil {
			log.Println(err, ";do you call client.SetHttpAgent()?")
		}
	}
	return
}

// 注册服务实例
func (client *ServiceClient) RegisterServiceInstance(param vo.RegisterServiceInstanceParam) (success bool, err error) {
	if len(param.Ip) <= 0 {
		err = errors.New("[client.RegisterServiceInstance] param.Ip can not be empty")
	}
	if err == nil && (param.Port <= 0 || param.Port > 65535) {
		err = errors.New("[client.RegisterServiceInstance] param.Port invalid")
	}
	if err == nil && len(param.ServiceName) <= 0 {
		err = errors.New("[client.RegisterServiceInstance] param.ServiceName can not be empty")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) + constant.SERVICE_PATH
		body := util.TransformObject2Param(param)
		header := map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		log.Println("[client.RegisterServiceInstance] request url:", path, " ;body:", body, " ;header:", header)
		responseTmp, errPost := agent.Post(path, header, clientConfig.TimeoutMs, body)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "ok" {
					success = true
				} else {
					success = false
					err = errors.New("[client.RegisterServiceInstance] " + string(bytes))
				}
			} else {
				err = errors.New("[client.RegisterServiceInstance] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

// 注销服务实例
func (client *ServiceClient) LogoutServiceInstance(param vo.LogoutServiceInstanceParam) (success bool, err error) {
	if len(param.ServiceName) <= 0 {
		err = errors.New("[client.LogoutServiceInstance] param.ServiceName can not be empty")
	}
	if err == nil && len(param.Ip) <= 0 {
		err = errors.New("[client.LogoutServiceInstance] param.Ip can not be empty")
	}
	if err == nil && (param.Port <= 0 || param.Port > 65535) {
		err = errors.New("[client.LogoutServiceInstance] param.Port invalid")
	}
	if err == nil && len(param.Cluster) <= 0 {
		err = errors.New("[client.LogoutServiceInstance] param.Cluster can not be empty")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) + constant.SERVICE_PATH
		params := util.TransformObject2Param(param)
		log.Println("[client.LogoutServiceInstance] request url:", path, ",params:", params)
		responseTmp, errPost := agent.Delete(path, nil, clientConfig.TimeoutMs, params)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "ok" {
					success = true
				} else {
					success = false
					err = errors.New("[client.LogoutServiceInstance] " + string(bytes))
				}
			} else {
				err = errors.New("[client.LogoutServiceInstance] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

// 修改服务实例
func (client *ServiceClient) ModifyServiceInstance(param vo.ModifyServiceInstanceParam) (success bool, err error) {
	if len(param.Ip) <= 0 {
		err = errors.New("[client.ModifyServiceInstance] param.Ip can not be empty")
	}
	if err == nil && (param.Port <= 0 || param.Port > 65535) {
		err = errors.New("[client.ModifyServiceInstance] param.Port invalid")
	}
	if err == nil && len(param.ServiceName) <= 0 {
		err = errors.New("[client.ModifyServiceInstance] param.ServiceName can not be empty")
	}
	if err == nil && len(param.Cluster) <= 0 {
		err = errors.New("[client.ModifyServiceInstance] param.Cluster can not be empty")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) + constant.SERVICE_PATH + "/update"
		body := util.TransformObject2Param(param)
		header := map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		log.Println("[client.ModifyServiceInstance] request url:", path, " ;body:", body, " ;header:", header)
		responseTmp, errPost := agent.Put(path, header, clientConfig.TimeoutMs, body)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "ok" {
					success = true
				} else {
					success = false
					err = errors.New("[client.ModifyServiceInstance] " + string(bytes))
				}
			} else {
				err = errors.New("[client.ModifyServiceInstance] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

// 获取服务列表
func (client *ServiceClient) GetService(param vo.GetServiceParam) (service vo.Service, err error) {
	if len(param.ServiceName) <= 0 {
		err = errors.New("[client.GetService] param.ServiceName can not be empty")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) + constant.SERVICE_PATH + "/list"
		params := util.TransformObject2Param(param)
		log.Println("[client.GetService] request url:", path,",params:",params)
		responseTmp, errPost := agent.Get(path, nil, clientConfig.TimeoutMs, params)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				errUnmarshal := json.Unmarshal(bytes, &service)
				if errUnmarshal != nil {
					log.Println(errUnmarshal)
					err = errors.New("[client.GetService] " + string(bytes))
				}
			} else {
				err = errors.New("[client.GetService] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

// 获取服务某个实例
func (client *ServiceClient) GetServiceInstance(param vo.GetServiceInstanceParam) (serviceInstance vo.ServiceInstance, err error) {
	if len(param.ServiceName) <= 0 {
		err = errors.New("[client.GetServiceInstance] param.ServiceName can not be empty")
	}
	if err == nil && len(param.Ip) <= 0 {
		err = errors.New("[client.GetServiceInstance] param.Ip can not be empty")
	}
	if err == nil && (param.Port <= 0 || param.Port > 65535) {
		err = errors.New("[client.GetServiceInstance] param.Port invalid")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) +
			constant.SERVICE_PATH
		params := util.TransformObject2Param(param)
		log.Println("[client.GetServiceInstance] request url:", path,",params:",params)
		responseTmp, errPost := agent.Get(path, nil, clientConfig.TimeoutMs, params)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				errUnmarshal := json.Unmarshal(bytes, &serviceInstance)
				if errUnmarshal != nil {
					log.Println(errUnmarshal)
					err = errors.New("[client.GetServiceInstance] " + string(bytes))
				}
			} else {
				err = errors.New("[client.GetServiceInstance] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

// 开始发送心跳的任务  只有在service.healthCheckMode = client的情况下才有效
func (client *ServiceClient) StartBeatTask(param vo.BeatTaskParam) (err error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.beating {
		err = errors.New("[client.StartBeatTask] client is beating,do not operator repeat")
	}
	// 开启任务
	if err == nil {
		client.beating = true
		client.startBeatTask(param)
	}
	return
}

func (client *ServiceClient) startBeatTask(param vo.BeatTaskParam) {
	go func() {
		for {
			clientConfig, serverConfigs, agent, errInner := client.sync()
			// 心跳参数检查
			if errInner == nil {
				if len(param.Ip) <= 0 {
					errInner = errors.New("[client.StartBeatTask] param.Ip can not be empty")
				}
				if errInner == nil && len(param.Dom) <= 0 {
					errInner = errors.New("[client.StartBeatTask] param.Dom can not be empty")
				}
			}
			if errInner != nil {
				client.mutex.Lock()
				client.beating = false
				client.mutex.Unlock()
				log.Println("client.StartBeatTask failed")
				break
			}
			// 检查service的健康检查模式
			if errInner == nil {
				serviceDetail, err := client.GetServiceDetail(vo.GetServiceDetailParam{
					ServiceName: param.Dom,
				})
				if err != nil {
					log.Println(err)
				}
				if serviceDetail.Service.HealthCheckMode != "client" {
					log.Println("[client.StartBeatTask] service.HealthCheckMode != 'client',sending a heartbeat is invalid")
				}
			}
			// 创建计时器
			var timer *time.Timer
			if errInner == nil {
				timer = time.NewTimer(time.Duration(clientConfig.BeatInterval) * time.Millisecond)
			}
			// http 请求
			if errInner == nil {
				path := client.buildBasePath(serverConfigs[0]) + constant.SERVICE_BASE_PATH + "/api/clientBeat"
				body := make(map[string]string)
				body[constant.KEY_DOM] = param.Dom
				paramBytes, errMarshal := json.Marshal(param)
				if errMarshal != nil {
					log.Println(errMarshal)
					continue
				}
				body[constant.KEY_BEAT] = string(paramBytes)
				header := map[string][]string{
					"Content-Type": {"application/x-www-form-urlencoded"},
				}
				log.Println("[client.StartBeatTask] request url:", path, " ;body:", body, " ;header:", header)
				response, errPost := agent.Post(path, header, clientConfig.TimeoutMs, body)
				if errPost != nil {
					log.Println(errPost)
					continue
				}
				bytes, errRead := ioutil.ReadAll(response.Body)
				if errRead != nil {
					log.Println(errRead)
					continue
				} else {
					if response.StatusCode == 200 {
						log.Print("[client.StartBeatTask] send beat success:" + string(bytes))
					} else {
						log.Println("[client.StartBeatTask] send beat failed:[" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
					}
					_ = response.Body.Close()
				}
			}
			if !client.beating {
				break
			}
			<-timer.C
		}
	}()
}

// 停止发送心跳的任务
func (client *ServiceClient) StopBeatTask() {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	client.beating = false
	log.Println("[client.StopBeatTask] client stop beating success")
}

func (client *ServiceClient) GetServiceDetail(param vo.GetServiceDetailParam) (serviceDetail vo.ServiceDetail, err error) {
	if len(param.ServiceName) <= 0 {
		err = errors.New("[client.GetServiceInfo] param.ServiceName can not be empty")
	}
	var clientConfig constant.ClientConfig
	var serverConfigs []constant.ServerConfig
	var agent http_agent.IHttpAgent
	if err == nil {
		clientConfig, serverConfigs, agent, err = client.sync()
	}
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := client.buildBasePath(serverConfigs[0]) +
			constant.SERVICE_BASE_PATH + "/catalog/serviceDetail"
		params := util.TransformObject2Param(param)
		log.Println("[client.GetServiceInfo] request url:", path, ",params:", params)
		responseTmp, errPost := agent.Get(path, nil, clientConfig.TimeoutMs, params)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	// response 解析
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				errUnmarshal := json.Unmarshal(bytes, &serviceDetail)
				if errUnmarshal != nil {
					log.Println(errUnmarshal)
					err = errors.New("[client.GetServiceInfo] " + string(bytes))
				}
			} else {
				err = errors.New("[client.GetServiceInfo] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

func (client *ServiceClient) buildBasePath(serverConfig constant.ServerConfig) (basePath string) {
	basePath = "http://" + serverConfig.IpAddr + ":" +
		strconv.FormatUint(serverConfig.Port, 10) + serverConfig.ContextPath + constant.CONFIG_PATH
	return
}
