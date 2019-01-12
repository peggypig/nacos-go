package http_agent

import "net/http"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-10 11:26
**/
type HttpAgent struct {
}

func (agent *HttpAgent) Get(path string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return get(path, header, timeoutMs, params)
}

func (agent *HttpAgent) Post(path string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return post(path, header, timeoutMs, params)
}
func (agent *HttpAgent) Delete(path string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return delete(path, header, timeoutMs, params)
}
func (agent *HttpAgent) Put(path string, header http.Header, timeoutMs uint64,
	params map[string]string) (response *http.Response, err error) {
	return put(path, header, timeoutMs, params)
}
