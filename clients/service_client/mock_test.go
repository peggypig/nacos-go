package service_client

import (
	"github.com/golang/mock/gomock"
	"nacos-go/clients/nacos_client"
	"nacos-go/common/constant"
	"nacos-go/common/http_agent"
	"nacos-go/vo"
	"net/http"
	"testing"
	"time"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-10 15:55
**/

func TestMockIServiceClient_GetService(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

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
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/list?serviceName=DEMO&clusters=a"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).Times(1).
		Return(http_agent.FakeHttpResponse(200, `{
	"dom": "DEMO",
	"cacheMillis": 1000,
	"useSpecifiedURL": false,
	"hosts": [{
		"valid": true,
		"marked": false,
		"instanceId": "10.10.10.10-8888-a-DEMO",
		"port": 8888,
		"ip": "10.10.10.10",
		"weight": 1.0,
		"metadata": {}
	}],
	"checksum": "3bbcf6dd1175203a8afdade0e77a27cd1528787794594",
	"lastRefTime": 1528787794594,
	"env": "",
	"clusters": ""
}`), nil)

	client := ServiceClient{}
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
	service, err := client.GetService(vo.GetServiceParam{
		ServiceName: "DEMO",
		Clusters:    []string{"a"},
	})
	t.Log(service, err)
}

func TestMockIServiceClient_GetServiceDetail(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

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
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail?serviceName=DEMO"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).Times(1).
		Return(http_agent.FakeHttpResponse(200, `{
		"service":{
			"name":"DEMO",
			"protectThreshold":0.0,
			"app":null,
			"group":null,
			"healthCheckMode":"client",
			"metadata":{}
		},
		"clusters":[]}`), nil)

	client := ServiceClient{}
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
	service, err := client.GetServiceDetail(vo.GetServiceDetailParam{
		ServiceName: "DEMO",
	})
	t.Log(service, err)
}

func TestMockIServiceClient_GetServiceInstance(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

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
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance?serviceName=DEMO&ip=10.10.10.10&port=80&healthyOnlyfalse"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).Times(1).
		Return(http_agent.FakeHttpResponse(200, `{
	"metadata": {},
	"instanceId": "10.10.10.10-8888-DEFAULT-DEMO",
	"port": 8888,
	"service": "DEMO",
	"healthy": false,
	"ip": "10.10.10.10",
	"clusterName": "DEFAULT",
	"weight": 1.0
}`), nil)

	client := ServiceClient{}
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
	service, err := client.GetServiceInstance(vo.GetServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.10.10.10",
		Port:        80,
	})
	t.Log(service, err)
}

func TestMockIServiceClient_RegisterServiceInstance(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"enable":      "false",
			"healthy":     "false",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
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
	service, err := client.RegisterServiceInstance(vo.RegisterServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
	})
	t.Log(service, err)
}

func TestMockIServiceClient_ModifyServiceInstance(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Put(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance/update"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"serviceName": "DEMO",
			"ip":          "10.0.0.10",
			"port":        "80",
			"weight":      "0",
			"cluster":     "DEFAULT",
		})).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
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
	service, err := client.ModifyServiceInstance(vo.ModifyServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	t.Log(service, err)
}

func TestMockIServiceClient_LogoutServiceInstance(t *testing.T) {
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
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/instance?serviceName=DEMO&ip=10.0.0.10&port=80&cluster=DEFAULT"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).Times(1).
		Return(http_agent.FakeHttpResponse(200, `ok`), nil)

	client := ServiceClient{}
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
	service, err := client.LogoutServiceInstance(vo.LogoutServiceInstanceParam{
		ServiceName: "DEMO",
		Ip:          "10.0.0.10",
		Port:        80,
		Cluster:     "DEFAULT",
	})
	t.Log(service, err)
}

func TestMockIServiceClient_StartBeatTask(t *testing.T) {
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
		BeatInterval:   5 * 1000,
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().AnyTimes().Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().AnyTimes().Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().AnyTimes().Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail?serviceName=DEMO"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, `{"service":
{	"name":"DEMO",
	"protectThreshold":0.0,
	"app":null,
	"group":null,
	"healthCheckMode":"client",
	"metadata":{}
},
"clusters":[]}`), nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/api/clientBeat"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dom":  "DEMO",
			"beat": `{"ip":"10.0.0.10","port":80,"weight":0,"dom":"DEMO","cluster":"DEFAULT","metaData":null}`,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, `beat`), nil)

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	err := client.StartBeatTask(vo.BeatTaskParam{
		Dom:     "DEMO",
		Ip:      "10.0.0.10",
		Port:    80,
		Cluster: "DEFAULT",
	})
	t.Log(err)
	time.Sleep(100 * time.Second)
}

func TestMockIServiceClient_StopBeatTask(t *testing.T) {
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
		BeatInterval:   5 * 1000,
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().AnyTimes().Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().AnyTimes().Return([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().AnyTimes().Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/catalog/serviceDetail?serviceName=DEMO"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000))).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, `{"service":
{	"name":"DEMO",
	"protectThreshold":0.0,
	"app":null,
	"group":null,
	"healthCheckMode":"client",
	"metadata":{}
},
"clusters":[]}`), nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/ns/api/clientBeat"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dom":  "DEMO",
			"beat": `{"ip":"10.0.0.10","port":80,"weight":0,"dom":"DEMO","cluster":"DEFAULT","metaData":null}`,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, `beat`), nil)

	client := ServiceClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   5 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr: "console.nacos.io",
		Port:   80,
	}})
	err := client.StartBeatTask(vo.BeatTaskParam{
		Dom:     "DEMO",
		Ip:      "10.0.0.10",
		Port:    80,
		Cluster: "DEFAULT",
	})
	go func() {
		time.Sleep(6*time.Second)
		client.StopBeatTask()
	}()
	t.Log(err)
	time.Sleep(21 * time.Second)
}
