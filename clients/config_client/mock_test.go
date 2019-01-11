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

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs?dataId=TEST&group=TEST"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	t.Log(content, err)
}
