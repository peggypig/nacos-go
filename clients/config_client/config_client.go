package config_client

import (
	"errors"
	"io/ioutil"
	"log"
	"nacos-go/common/constant"
	"nacos-go/common/httpagent"
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
* @create : 2019-01-08 10:01
**/

type ConfigClient struct {
	ServerConfigs []constant.ServerConfig
	ClientConfig  constant.ClientConfig
	localConfigs  []vo.ConfigParam
	mutex         sync.Mutex
	listening     bool
}

func (client *ConfigClient) GetConfigContent(dataId string, group string) (content string, err error) {
	if len(dataId) <= 0 {
		err = errors.New("[client.GetConfigContent] dataId param can not be empty")
	}
	if err == nil && len(group) <= 0 {
		err = errors.New("[client.GetConfigContent] group param can not be empty")
	}
	if err == nil {
		client.mutex.Lock()
		defer client.mutex.Unlock()
		exist := false
		for _, config := range client.localConfigs {
			if config.Group == group && config.DataId == dataId {
				content = config.Content
				exist = true
				break
			}
		}
		if !exist || len(content) <= 0 {
			content, err = client.GetConfig(vo.ConfigParam{
				DataId: dataId,
				Group:  group,
			})
		}
	}
	return
}

func (client *ConfigClient) GetConfig(param vo.ConfigParam) (content string, err error) {
	if len(param.DataId) <= 0 {
		err = errors.New("[client.GetConfig]=>param.dataId can not be empty")
	}
	if len(param.Group) <= 0 {
		err = errors.New("[client.GetConfig]=>param.group can not be empty")
	}
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
			strconv.FormatUint(client.ServerConfigs[0].Port, 10) + constant.CONFIG_PATH + "?dataId=" + param.DataId +
			"&group=" + param.Group
		if len(param.Tenant) > 0 {
			path += "&tenant=" + param.Tenant
		}
		log.Println("[client.GetConfig] request url :", path)
		responseTmp, errGet := httpagent.Get(path, nil, client.ClientConfig.TimeoutMs)
		if errGet != nil {
			err = errGet
		} else {
			response = responseTmp
		}
	}
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				content = string(bytes)
			} else {
				err = errors.New("[" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

func (client *ConfigClient) PublishConfig(param vo.ConfigParam) (published bool, err error) {
	if len(param.DataId) <= 0 {
		err = errors.New("[client.PublishConfig]=>param.dataId can not be empty")
	}
	if len(param.Group) <= 0 {
		err = errors.New("[client.PublishConfig]=>param.group can not be empty")
	}
	if len(param.Content) <= 0 {
		err = errors.New("[client.PublishConfig]=>param.content can not be empty")
	}
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
			strconv.FormatUint(client.ServerConfigs[0].Port, 10) + constant.CONFIG_PATH
		body := make(map[string]string)
		body[constant.KEY_DATA_ID] = param.DataId
		body[constant.KEY_GROUP] = param.Group
		body[constant.KEY_CONTENT] = param.Content
		if len(param.Tenant) > 0 {
			body[constant.KEY_TENANT] = param.Tenant
		}
		if len(param.Desc) > 0 {
			body[constant.KEY_DESC] = param.Desc
		}
		if len(param.AppName) > 0 {
			body[constant.KEY_APP_NAME] = param.AppName
		}
		header := map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		log.Println("[client.PublishConfig] request url:", path, " ;body:", body, " ;header:", header)
		responseTmp, errPost := httpagent.Post(path, header, client.ClientConfig.TimeoutMs, body)
		if errPost != nil {
			err = errPost
		} else {
			response = responseTmp
		}
	}
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "true" {
					published = true
				} else {
					published = false
					err = errors.New(string(bytes))
				}
			} else {
				err = errors.New("[" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

func (client *ConfigClient) DeleteConfig(param vo.ConfigParam) (deleted bool, err error) {
	if len(param.DataId) <= 0 {
		err = errors.New("[client.DeleteConfig]=>param.dataId can not be empty")
	}
	if len(param.Group) <= 0 {
		err = errors.New("[client.DeleteConfig]=>param.group can not be empty")
	}
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
			strconv.FormatUint(client.ServerConfigs[0].Port, 10) + constant.CONFIG_PATH + "?dataId=" + param.DataId +
			"&group=" + param.Group
		if len(param.Tenant) > 0 {
			path += "&tenant=" + param.Tenant
		}
		log.Println("[client.DeleteConfig] request url:", path)
		responseTmp, errDelete := httpagent.Delete(path, nil, client.ClientConfig.TimeoutMs)
		if errDelete != nil {
			err = errDelete
		} else {
			response = responseTmp
		}
	}
	if err == nil {
		bytes, errRead := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if errRead != nil {
			err = errRead
		} else {
			if response.StatusCode == 200 {
				if strings.ToLower(strings.Trim(string(bytes), " ")) == "true" {
					deleted = true
				} else {
					deleted = false
					err = errors.New(string(bytes))
				}
			} else {
				err = errors.New("[" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
			}
		}
	}
	return
}

func (client *ConfigClient) ListenConfig(params []vo.ConfigParam) (err error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.listening {
		err = errors.New("client is listening,do not operator repeat")
	}
	// 监听
	if err == nil {
		client.localConfigs = params
		client.listenTask()
	}
	return
}

func (client *ConfigClient) listenTask() {
	go func() {
		for {
			timer := time.NewTimer(time.Duration(client.ClientConfig.ListenInterval) * time.Millisecond)
			var listeningConfigs string
			var errInner error
			// 检查&拼接监听参数
			for index, param := range client.localConfigs {
				if len(param.DataId) <= 0 {
					errInner = errors.New("[client.ListenConfig] params[" + strconv.Itoa(index) + "].DataId can not be empty")
					break
				}
				if len(param.Group) <= 0 {
					errInner = errors.New("[client.ListenConfig] params[" + strconv.Itoa(index) + "].Group can not be empty")
					break
				}
				var tenant string
				if len(param.Tenant) > 0 {
					tenant = param.Tenant
				}
				var md5 string
				if len(param.Content) > 0 {
					md5 = util.Md5(param.Content)
				}
				listeningConfigs += param.DataId + constant.SPLIT_CONFIG_INNER + param.Group + constant.SPLIT_CONFIG_INNER +
					md5 + constant.SPLIT_CONFIG_INNER + tenant + constant.SPLIT_CONFIG
			}
			if errInner != nil {
				log.Println(errInner, ";client.ListenConfig failed")
				break
			}
			if errInner == nil {
				// http 请求
				path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
					strconv.FormatUint(client.ServerConfigs[0].Port, 10) + constant.CONFIG_LISTEN_PATH
				body := make(map[string]string)
				body[constant.KEY_LISTEN_CONFIGS] = listeningConfigs
				header := map[string][]string{
					"Content-Type":         {"application/x-www-form-urlencoded"},
					"Long-Pulling-Timeout": {strconv.FormatUint(client.ClientConfig.ListenInterval, 10)},
				}
				log.Println("[client.ListenConfig] request url:", path, " ;body:", body, " ;header:", header)
				response, errPost := httpagent.Post(path, header, client.ClientConfig.TimeoutMs, body)
				if errPost != nil {
					log.Println(errPost)
					continue
				}
				var changed string
				bytes, errRead := ioutil.ReadAll(response.Body)
				if errRead != nil {
					log.Println(errRead)
					continue
				} else {
					if response.StatusCode == 200 {
						if strings.ToLower(strings.Trim(string(bytes), " ")) == "" {
							log.Println("[client.ListenConfig] no change")
						} else {
							changed = string(bytes)
							log.Print("[client.ListenConfig] config changed:" + changed)
						}
					} else {
						log.Println("[" + strconv.Itoa(response.StatusCode) + "]" + string(bytes))
					}
					_ = response.Body.Close()
				}
				if len(changed) > 0 {
					client.updateLocalConfig(changed)
				}
				client.listening = true
				if !client.listening {
					break
				}
			}
			<-timer.C
		}
	}()
}

func (client *ConfigClient) StopListenConfig() {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	if client.listening {
		client.listening = false
	}
	log.Println("[client.StopListenConfig] stop listen config success")
}

// ListenConfig 发现配置变化时候，修改本地配置
func (client *ConfigClient) updateLocalConfig(changed string) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	changedConfigs := strings.Split(changed, "%01")
	for _, config := range changedConfigs {
		attrs := strings.Split(config, "%02")
		if len(attrs) == 2 {
			content, err := client.GetConfig(vo.ConfigParam{
				DataId: attrs[0],
				Group:  attrs[1],
			})
			if err != nil {
				log.Println("[client.updateLocalConfig] update config failed:", err.Error())
			} else {
				client.putLocalConfig(vo.ConfigParam{
					DataId:  attrs[0],
					Group:   attrs[1],
					Content: content,
				})
			}
		}
	}
	log.Println("[client.updateLocalConfig] update config complete")
	log.Println("[client.localConfig]:", client.localConfigs)
}

func (client *ConfigClient) putLocalConfig(config vo.ConfigParam) {
	if len(config.DataId) > 0 && len(config.Group) > 0 {
		exist := false
		for i := 0; i < len(client.localConfigs); i++ {
			if len(client.localConfigs[i].DataId) > 0 && len(client.localConfigs[i].Group) > 0 &&
				config.DataId == client.localConfigs[i].DataId && config.Group == client.localConfigs[i].Group {
				// 本地存在 则更新
				client.localConfigs[i] = config
				exist = true
				break
			}
		}
		if !exist {
			// 本地不存在 放入
			client.localConfigs = append(client.localConfigs, config)
		}
	}
	log.Println("[client.putLocalConfig] putLocalConfig success")
}
