package util

import (
	"fmt"
	"github.com/peggypig/nacos-go/vo"
	"testing"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-11 19:25
**/

func TestTransformObject2Param(t *testing.T) {
	object := vo.GetServiceParam{
		Tenant:      "aaa",
		HealthyOnly: true,
	}
	params := TransformObject2Param(object)
	fmt.Println(params)
}
