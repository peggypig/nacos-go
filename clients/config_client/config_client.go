package config_client

import (
	"log"
	"nacos-go/common/http"
	"nacos-go/common/util"
	"nacos-go/vo"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

type ConfigClient struct {
	PostTimeout              int64
	Agent                    http.HttpAgent
	NameSpace                string
	Encode                   string
	Worker                   ClientWorker
	configFilterChainManager vo.ConfigFilterChainManager
}

func (client *ConfigClient) AddListener() {

}

func (client *ConfigClient) RemoveListener() {

}

func (client *ConfigClient) GetConfig(dataId string, group string, timeoutMS int64) (content string) {
	err := util.CheckDataId(dataId)
	if err != nil {
		panic(err.Error())
	}
	err = util.CheckGroup(group)
	if err != nil {
		panic(err.Error())
	}
	var response = &vo.ConfigResponse{}
	response.SetDataId(dataId)
	response.SetGroup(group)
	response.SetTenant(client.NameSpace)

	// 优先使用本地配置
	contentTemp := getFailover(client.Agent.GetName(), dataId, group, response.GetTenant())
	if contentTemp != nil {
		response.SetContent(contentTemp.(string))
		client.configFilterChainManager.DoFilter(nil, response)
		content = response.GetContent()
	} else {

	}
	return
}

func (client *ConfigClient) PublishConfig() {

}

func (client *ConfigClient) RemoveConfig() {

}

func (client *ConfigClient) GetServerStatus() (status string) {

}
