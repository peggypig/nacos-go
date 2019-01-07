package http

import "nacos-go/vo"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:36
**/

type HttpAgent struct {
	Properties map[string]string
	Slm        vo.ServerListManager
}

func (agent *HttpAgent) Start() {

}

func (agent *HttpAgent) GetName() string {
	return agent.Slm.Name
}
