package config_client

import (
	"errors"
	"io/ioutil"
	"log"
	"nacos-go/common/constant"
	"nacos-go/common/httpagent"
	"nacos-go/vo"
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
* @create : 2019-01-08 10:01
**/

type ConfigClient struct {
	Tenant        string
	ServerConfigs []constant.ServerConfig
	ClientConfig  constant.ClientConfig
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
			strconv.FormatUint(uint64(client.ServerConfigs[0].Port), 10) + constant.CONFIG_BASE_PATH + "?dataId=" + param.DataId +
			"&group=" + param.Group
		if len(client.Tenant) > 0 {
			path += "&tenant=" + client.Tenant
		}
		log.Println("[client.GetConfig] request url :", path)
		responseTmp, errGet := httpagent.Get(path, nil, uint64(client.ClientConfig.TimeoutMs))
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
		err = errors.New("[client.GetConfig]=>param.dataId can not be empty")
	}
	if len(param.Group) <= 0 {
		err = errors.New("[client.GetConfig]=>param.group can not be empty")
	}
	if len(param.Content) <= 0 {
		err = errors.New("[client.GetConfig]=>param.content can not be empty")
	}
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
			strconv.FormatUint(uint64(client.ServerConfigs[0].Port), 10) + constant.CONFIG_BASE_PATH
		body := make(map[string]string)
		body[constant.KEY_DATA_ID] = param.DataId
		body[constant.KEY_GROUP] = param.Group
		body[constant.KEY_CONTENT] = param.Content
		if len(client.Tenant) > 0 {
			body[constant.KEY_TENANT] = param.Tenant
		}
		if len(param.Desc) > 0 {
			body[constant.KEY_DESC] = param.Desc
		}
		if len(param.AppName) > 0 {
			body[constant.KEY_APP_NAME] = param.AppName
		}
		log.Println("[client.GetConfig] request url:", path, " ;body:", body)
		responseTmp, errPost := httpagent.Post(path, map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}, uint64(client.ClientConfig.TimeoutMs), body)
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
	return
}
func (client *ConfigClient) ListenConfig() {

}
