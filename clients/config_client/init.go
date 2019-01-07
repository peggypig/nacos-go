package config_client

import (
	"log"
	"os"
	"path/filepath"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 17:06
**/

func init() {
	localFileRootPath = os.Getenv("JM.LOG.PATH")
	if len(localFileRootPath) <= 0 {
		localFileRootPath = os.Getenv("USER") + os.Getenv("HOME") +
			string(filepath.Separator) + "nacos" + string(filepath.Separator) + "config"
	}
	localSnapshotPath = os.Getenv("JM.SNAPSHOT.PATH")
	if len(localSnapshotPath) <= 0 {
		localSnapshotPath = os.Getenv("USER") + os.Getenv("HOME") +
			string(filepath.Separator) + "nacos" + string(filepath.Separator) + "config"
	}
	log.Println("LocalFileRootPath:", localFileRootPath, "\tLocalSnapshotPath:", localSnapshotPath)
}
