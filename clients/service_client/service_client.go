package service_client

import (
	"encoding/json"
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
* @create : 2019-01-07 15:13
**/

type ServiceClient struct {
	ServerConfigs []constant.ServerConfig
	ClientConfig  constant.ClientConfig
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
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" + strconv.FormatUint(client.ServerConfigs[0].Port, 10) +
			constant.SERVICE_PATH
		body := make(map[string]string)
		body[constant.KEY_SERVICE_NAME] = param.ServiceName
		body[constant.KEY_IP] = param.Ip
		body[constant.KEY_PORT] = strconv.FormatUint(param.Port, 10)
		body[constant.KEY_ENABLE] = strconv.FormatBool(param.Enable)
		body[constant.KEY_HEALTHY] = strconv.FormatBool(param.Healthy)
		if len(param.Tenant) > 0 {
			body[constant.KEY_TENANT] = param.Tenant
		}
		if len(param.ClusterName) > 0 {
			body[constant.KEY_CLUSTER_NAME] = param.ClusterName
		}
		if param.Weight >= 0 {
			body[constant.KEY_WEIGHT] = strconv.FormatFloat(param.Weight, 'f', -1, 64)
		}
		if len(param.Metadata) > 0 {
			body[constant.KEY_METADATA] = param.Metadata
		}
		header := map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		log.Println("[client.RegisterServiceInstance] request url:", path, " ;body:", body, " ;header:", header)
		responseTmp, errPost := httpagent.Post(path, header, client.ClientConfig.TimeoutMs, body)
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
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" + strconv.FormatUint(client.ServerConfigs[0].Port, 10) +
			constant.SERVICE_PATH + "?serviceName=" + param.ServiceName + "&ip=" + param.Ip + "&port=" +
			strconv.FormatUint(param.Port, 10)
		if len(param.Tenant) > 0 {
			path += "&tenant=" + param.Tenant
		}
		log.Println("[client.LogoutServiceInstance] request url:", path)
		responseTmp, errPost := httpagent.Delete(path, nil, client.ClientConfig.TimeoutMs)
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
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" + strconv.FormatUint(client.ServerConfigs[0].Port, 10) +
			constant.SERVICE_PATH + "/update"
		body := make(map[string]string)
		body[constant.KEY_SERVICE_NAME] = param.ServiceName
		body[constant.KEY_IP] = param.Ip
		body[constant.KEY_PORT] = strconv.FormatUint(param.Port, 10)
		if len(param.Tenant) > 0 {
			body[constant.KEY_TENANT] = param.Tenant
		}
		if param.Weight >= 0 {
			body[constant.KEY_WEIGHT] = strconv.FormatFloat(param.Weight, 'f', -1, 64)
		}
		if len(param.Metadata) > 0 {
			body[constant.KEY_METADATA] = param.Metadata
		}
		header := map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
		log.Println("[client.ModifyServiceInstance] request url:", path, " ;body:", body, " ;header:", header)
		responseTmp, errPost := httpagent.Put(path, header, client.ClientConfig.TimeoutMs, body)
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
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" + strconv.FormatUint(client.ServerConfigs[0].Port, 10) +
			constant.SERVICE_PATH + "/list?serviceName=" + param.ServiceName
		if len(param.Tenant) > 0 {
			path += "&tenant=" + param.Tenant
		}
		if len(param.Clusters) > 0 {
			clusters := ""
			for _, cluster := range param.Clusters {
				if len(cluster) > 0 {
					clusters += cluster + ","
				}
			}
			if strings.HasSuffix(clusters, ",") {
				clusters = clusters[:len(clusters)-1]
			}
			if len(clusters) > 0 {
				path += "&clusters=" + clusters
			}
		}
		log.Println("[client.GetService] request url:", path)
		responseTmp, errPost := httpagent.Get(path, nil, client.ClientConfig.TimeoutMs)
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
	// 构造并完成http请求
	var response *http.Response
	if err == nil {
		path := "http://" + client.ServerConfigs[0].IpAddr + ":" + strconv.FormatUint(client.ServerConfigs[0].Port, 10) +
			constant.SERVICE_PATH + "?serviceName=" + param.ServiceName + "&ip=" + param.Ip + "&port=" +
			strconv.FormatUint(param.Port, 10) + "&healthyOnly" + strconv.FormatBool(param.HealthyOnly)
		if len(param.Tenant) > 0 {
			path += "&tenant=" + param.Tenant
		}
		if len(param.Cluster) > 0 {
			path += "&cluster=" + param.Cluster
		}
		log.Println("[client.GetServiceInstance] request url:", path)
		responseTmp, errPost := httpagent.Get(path, nil, client.ClientConfig.TimeoutMs)
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
