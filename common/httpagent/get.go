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
	endTime := time.Now().UnixNano()/1e9 + int64(timeoutMs)
	for time.Now().UnixNano()/1e9 <= endTime {
		client := http.Client{}
		request, errNew := http.NewRequest("GET", path, nil)
		if errNew != nil {
			log.Println(errNew)
			err = errNew
			break
		}
		request.Header = header
		resp, errDo := client.Do(request)
		if errDo != nil {
			log.Println(errDo)
			err = errDo
			break
		}
		response = resp
		err = nil
	}
	return
}
