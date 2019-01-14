package http_agent

import (
	"io/ioutil"
	"testing"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-14 10:25
**/

func TestHttpAgent_Get(t *testing.T) {
	agent := HttpAgent{}
	response, err := agent.Get("http://10.0.0.8:88148/nacos1/v1/cs/configs", nil, 10000, map[string]string{
		"dataId": "TEST21",
		"group":  "DEFAULT_GROUP",
	})
	if err != nil {
		t.Error(err)
	} else {
		bytes, _ := ioutil.ReadAll(response.Body)
		t.Log("Body:", string(bytes))
		t.Log("Status:", response.Status)
	}
}
