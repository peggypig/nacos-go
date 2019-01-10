package http_agent

import "net/http"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-10 11:07
**/

//go:generate mockgen -destination mock_http_agent_interface.go -package http_agent nacos-go/common/http_agent IHttpAgent

type IHttpAgent interface {
	Get(path string, header http.Header, timeoutMs uint64) (response *http.Response, err error)
	Post(path string, header http.Header, timeoutMs uint64, params map[string]string) (response *http.Response, err error)
	Delete(path string, header http.Header, timeoutMs uint64) (response *http.Response, err error)
	Put(path string, header http.Header, timeoutMs uint64, params map[string]string) (response *http.Response, err error)
}
