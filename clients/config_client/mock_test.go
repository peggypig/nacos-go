package config_client

import (
	"github.com/golang/mock/gomock"
	"nacos-go/clients/nacos_client"
	"nacos-go/common/constant"
	"nacos-go/common/http_agent"
	"nacos-go/vo"
	"net/http"
	"testing"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-10 15:55
**/


func TestMockIConfigClient_GetConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	callSetHttpAgent := mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	//mockINacosClient.EXPECT()

	callGetClientConfig :=mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{},nil).
		After(callSetHttpAgent)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{},nil).
		After(callGetClientConfig)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1)

	var resp *http.Response
	mockIHttpAgent.EXPECT().Get(gomock.Eq("http://console.nacos.io/nacos/v1/cs/configs?dataId=TEST&group=TEST"), gomock.Nil(),
		gomock.Eq(30*1000)).Times(1).Return(resp, nil)


	client := &ConfigClient{}
	client.INacosClient = mockINacosClient
	client.SetHttpAgent(mockIHttpAgent)
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	t.Log(content,err)
}
