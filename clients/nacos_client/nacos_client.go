package nacos_client

import (
	"encoding/json"
	"errors"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/common/nacos_error"
	"github.com/peggypig/nacos-go/common/util"
	"github.com/peggypig/nacos-go/vo"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-09 16:39
**/

type NacosClient struct {
	clientConfigValid  bool
	serverConfigsValid bool
	agent              http_agent.IHttpAgent
	clientConfig       constant.ClientConfig
	serverConfigs      []constant.ServerConfig
}

// 设置 clientConfig
func (client *NacosClient) SetClientConfig(config constant.ClientConfig) (err error) {
	if config.TimeoutMs <= 0 {
		err = errors.New("[client.SetClientConfig] config.TimeoutMs should > 0")
	}
	if err == nil && config.TimeoutMs < config.ListenInterval {
		err = errors.New("[client.SetClientConfig] config.TimeoutMs should >= config.ListenInterval")
	}
	if err == nil {
		if config.BeatInterval <= 0 {
			config.BeatInterval = 5 * 1000
		}
		if config.ListenInterval < 10*1000 {
			config.ListenInterval = 10 * 1000
		}
	}
	if err == nil {
		client.clientConfig = config
		client.clientConfigValid = true
	}
	return
}

// 设置 serverConfigs
func (client *NacosClient) SetServerConfig(configs []constant.ServerConfig) (err error) {
	if len(configs) <= 0 {
		err = errors.New("[client.SetServerConfig] configs can not be empty")
	}
	if err == nil {
		for i := 0; i < len(configs); i++ {
			if len(configs[i].IpAddr) <= 0 || configs[i].Port <= 0 || configs[i].Port > 65535 {
				err = errors.New("[client.SetServerConfig] configs[" + strconv.Itoa(i) + "] is invalid")
				break
			}
			if len(configs[i].ContextPath) <= 0 {
				configs[i].ContextPath = constant.DEFAULT_CONTEXT_PATH
			}
		}
	}
	if err == nil {
		client.serverConfigs = configs
		client.serverConfigsValid = true
	}
	return
}

// 获取 clientConfig
func (client *NacosClient) GetClientConfig() (config constant.ClientConfig, err error) {
	config = client.clientConfig
	if !client.clientConfigValid {
		err = errors.New("[client.GetClientConfig] invalid client config")
	}
	return
}

// 获取serverConfigs
func (client *NacosClient) GetServerConfig() (configs []constant.ServerConfig, err error) {
	configs = client.serverConfigs
	if !client.serverConfigsValid {
		err = errors.New("[client.GetServerConfig] invalid server configs")
	}
	return
}

func (client *NacosClient) SetHttpAgent(agent http_agent.IHttpAgent) (err error) {
	if agent == nil {
		err = errors.New("[client.SetHttpAgent] http agent can not be nil")
	} else {
		client.agent = agent
	}
	return
}

func (client *NacosClient) GetHttpAgent() (agent http_agent.IHttpAgent, err error) {
	if client.agent == nil {
		err = errors.New("[client.GetHttpAgent] invalid http agent")
	} else {
		agent = client.agent
	}
	return
}

func (client *NacosClient) check() (err error) {
	_, err = client.GetClientConfig()
	if err != nil {
		log.Println(err, ";do you call client.SetClientConfig()?")
	}
	if err == nil {
		_, err = client.GetServerConfig()
		if err != nil {
			log.Println(err, ";do you call client.SetServerConfig()?")
		}
	}
	if err == nil {
		_, err = client.GetHttpAgent()
		if err != nil {
			log.Println(err, ";do you call client.SetHttpAgent()?")
		}
	}
	return
}

func (client *NacosClient) GetNamespace() (namespaces []vo.Namespace, err error) {
	err = client.check()
	if err == nil {
		for _, serverConfig := range client.serverConfigs {
			path := client.buildBasePath(serverConfig)
			namespaces, err = getNamespace(client.agent, path, client.clientConfig.TimeoutMs, nil)
			if err == nil {
				break
			} else {
				if _, ok := err.(*nacos_error.NacosError); ok {
					break
				} else {
					log.Println("[client.GetNamespace] get namespace failed:", err.Error())
				}
			}
		}
	}
	return
}

func getNamespace(agent http_agent.IHttpAgent, path string, timeoutMs uint64,
	params map[string]string) (namespaces []vo.Namespace, err error) {
	var response *http.Response
	log.Println("[client.GetNamespace] request url :", path, ",params:", params)
	response, err = agent.Get(path, nil, timeoutMs, params)
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				var getNamespaceResponse vo.GetNamespaceResponse
				errUnmarshal := json.Unmarshal(bytes, &getNamespaceResponse)
				if errUnmarshal != nil {
					log.Println(errUnmarshal)
					err = errors.New("[client.GetNamespace] " + string(bytes))
				} else {
					if getNamespaceResponse.Code == 200 {
						namespaces = getNamespaceResponse.Data
					} else {
						err = errors.New("[client.GetNamespace] " + string(bytes))
					}
				}
			} else {
				err = &nacos_error.NacosError{
					ErrMsg: "[client.GetNamespace] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes),
				}
			}
		}
	}
	return
}

func (client *NacosClient) CreateNamespace(param vo.CreateNamespaceParam) (success bool, err error) {
	if len(param.NamespaceName) <= 0 {
		err = errors.New("[client.CreateNamespace] namespaceName param can not be empty")
	}
	if len(param.NamespaceDesc) <= 0 {
		err = errors.New("[client.CreateNamespace] namespaceDesc param can not be empty")
	}
	if err == nil {
		err = client.check()
	}
	if err == nil {
		params := util.TransformObject2Param(param)
		for _, serverConfig := range client.serverConfigs {
			path := client.buildBasePath(serverConfig)
			success, err = createNamespace(client.agent, path, client.clientConfig.TimeoutMs, params)
			if err == nil {
				break
			} else {
				if _, ok := err.(*nacos_error.NacosError); ok {
					break
				} else {
					log.Println("[client.CreateNamespace] create namespace failed:", err.Error())
				}
			}
		}
	}
	return
}

func createNamespace(agent http_agent.IHttpAgent, path string, timeoutMs uint64,
	params map[string]string) (success bool, err error) {
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	var response *http.Response
	log.Println("[client.CreateNamespace] request url :", path, ",params:", params, ",header:", header)
	response, err = agent.Post(path, header, timeoutMs, params)
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "true" {
					success = true
				} else {
					success = false
					err = errors.New("[client.CreateNamespace] " + string(bytes))
				}
			} else {
				err = &nacos_error.NacosError{
					ErrMsg: "[client.CreateNamespace] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes),
				}
			}
		}
	}
	return
}

func (client *NacosClient) ModifyNamespace(param vo.ModifyNamespaceParam) (success bool, err error) {
	if len(param.NamespaceName) <= 0 {
		err = errors.New("[client.ModifyNamespace] namespaceName param can not be empty")
	}
	if len(param.NamespaceDesc) <= 0 {
		err = errors.New("[client.ModifyNamespace] namespaceDesc param can not be empty")
	}
	if len(param.Namespace) <= 0 {
		err = errors.New("[client.ModifyNamespace] namespace param can not be empty")
	}
	if err == nil {
		err = client.check()
	}
	if err == nil {
		params := util.TransformObject2Param(param)
		for _, serverConfig := range client.serverConfigs {
			path := client.buildBasePath(serverConfig)
			success, err = modifyNamespace(client.agent, path, client.clientConfig.TimeoutMs, params)
			if err == nil {
				break
			} else {
				if _, ok := err.(*nacos_error.NacosError); ok {
					break
				} else {
					log.Println("[client.ModifyNamespace] create namespace failed:", err.Error())
				}
			}
		}
	}
	return
}

func modifyNamespace(agent http_agent.IHttpAgent, path string, timeoutMs uint64,
	params map[string]string) (success bool, err error) {
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	var response *http.Response
	log.Println("[client.ModifyNamespace] request url :", path, ",params:", params, ",header:", header)
	response, err = agent.Put(path, header, timeoutMs, params)
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "true" {
					success = true
				} else {
					success = false
					err = errors.New("[client.ModifyNamespace] " + string(bytes))
				}
			} else {
				err = &nacos_error.NacosError{
					ErrMsg: "[client.ModifyNamespace] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes),
				}
			}
		}
	}
	return
}

func (client *NacosClient) DeleteNamespace(param vo.DeleteNamespaceParam) (success bool, err error) {
	if len(param.NamespaceId) <= 0 {
		err = errors.New("[client.ModifyNamespace] namespaceId param can not be empty")
	}
	if err == nil {
		err = client.check()
	}
	if err == nil {
		params := util.TransformObject2Param(param)
		for _, serverConfig := range client.serverConfigs {
			path := client.buildBasePath(serverConfig)
			success, err = deleteNamespace(client.agent, path, client.clientConfig.TimeoutMs, params)
			if err == nil {
				break
			} else {
				if _, ok := err.(*nacos_error.NacosError); ok {
					break
				} else {
					log.Println("[client.ModifyNamespace] create namespace failed:", err.Error())
				}
			}
		}
	}
	return
}

func deleteNamespace(agent http_agent.IHttpAgent, path string, timeoutMs uint64,
	params map[string]string) (success bool, err error) {
	var response *http.Response
	log.Println("[client.DeleteNamespace] request url :", path, ",params:", params)
	response, err = agent.Delete(path, nil, timeoutMs, params)
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "true" {
					success = true
				} else {
					success = false
					err = errors.New("[client.DeleteNamespace] " + string(bytes))
				}
			} else {
				err = &nacos_error.NacosError{
					ErrMsg: "[client.DeleteNamespace] [" + strconv.Itoa(response.StatusCode) + "]" + string(bytes),
				}
			}
		}
	}
	return
}

func (client *NacosClient) buildBasePath(serverConfig constant.ServerConfig) (basePath string) {
	basePath = "http://" + serverConfig.IpAddr + ":" +
		strconv.FormatUint(serverConfig.Port, 10) + serverConfig.ContextPath + constant.NAMESPACE_PATH
	return
}
