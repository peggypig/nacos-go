package config_client

import (
	"errors"
	"io/ioutil"
	"log"
	"nacos-go/common/constant"
	"nacos-go/common/httpagent"
	"net/http"
	"strconv"
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

func (client *ConfigClient) GetConfig(dataId string, group string) (content string, err error) {
	if len(dataId) <= 0 {
		err = errors.New("[client.GetConfig]=>param dataId can not be empty")
	}
	if len(group) <= 0 {
		err = errors.New("[client.GetConfig]=>param group can not be empty")
	}
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" +
			strconv.FormatUint(uint64(client.ServerConfigs[0].Port), 10) + "/nacos/v1/cs/configs?dataId=" + dataId +
			"&group=" + group
		if len(client.Tenant) > 0{
			path +=  "&tenant=" + client.Tenant
		}
		log.Println("[client.GetConfig] request url :",path)
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
				err = errors.New(string(bytes))
			}
		}
	}
	return
}
func (client *ConfigClient) PublishConfig(dataId string, group string, content string) (published bool, err error) {
	return
}
func (client *ConfigClient) DeleteConfig(dataId string, group string) (deleted bool, err error) {
	return
}
func (client *ConfigClient) ListenConfig() {

}
