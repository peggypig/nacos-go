package config_client

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 17:02
**/

var localFileRootPath string
var localSnapshotPath string

func getFailover(serverName, dataId, group, tenant string) (failover interface{}) {
	fileInfo, filePath, err := getFailoverFile(serverName, dataId, group, tenant)
	if err == nil || !fileInfo.IsDir() {
		failover = readFile(filePath)
	}
	return
}

func getFailoverFile(serverName, dataId, group, tenant string) (fileInfo os.FileInfo, filePath string, err error) {
	filePath = localFileRootPath + string(filepath.Separator) + serverName + "_nacos"
	if len(tenant) <= 0 {
		filePath += string(filepath.Separator) + "config-data"
	} else {
		filePath += string(filepath.Separator) + "config-data-tenant" + string(filepath.Separator) + tenant
	}
	filePath += string(filepath.Separator) + group + string(filepath.Separator) + dataId
	fileInfo, err = os.Stat(filePath)
	return fileInfo, filePath, err
}

func readFile(filePath string) (content interface{}) {
	if strings.ToLower(strings.Trim(os.Getenv("isMultiInstance"), " ")) == "true" {
		// 多实例
	} else {
		file, errOpen := os.Open(filePath)
		if errOpen != nil {
			content = nil
			log.Println(errOpen)
		} else {
			defer file.Close()
			bytes, errRead := ioutil.ReadAll(file)
			if errRead != nil {
				content = nil
				log.Println(errRead)
			} else {
				content = string(bytes)
			}
		}
	}
	return
}
