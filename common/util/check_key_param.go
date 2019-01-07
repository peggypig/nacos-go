package util

import "errors"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 16:24
**/

func CheckDataId(dataId string) (err error) {
	if checkParamValue(dataId) != nil {
		err = errors.New("dataId can not be empty")
	}
	return
}

func CheckGroup(group string) (err error) {
	if checkParamValue(group) != nil {
		err = errors.New("group can not be empty")
	}
}

func checkParamValue(value string) (err error) {
	if len(value) <= 0 {
		err = errors.New("")
	}
	return
}
