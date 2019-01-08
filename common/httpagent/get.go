package httpagent

import (
	"log"
	"net/http"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

func Get(path string, header http.Header, timeoutMs uint64) (response *http.Response, err error) {
	client := http.Client{}
	client.Timeout = time.Millisecond * time.Duration(timeoutMs)
	request, errNew := http.NewRequest(http.MethodGet, path, nil)
	if errNew != nil {
		log.Println(errNew)
		err = errNew
	}
	if err == nil {
		request.Header = header
		resp, errDo := client.Do(request)
		if errDo != nil {
			log.Println(errDo)
			err = errDo
		} else {
			response = resp
		}
	}
	return
}

