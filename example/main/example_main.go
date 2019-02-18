package main

import (
	"github.com/peggypig/nacos-go/example"
	"time"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-02-18 10:52
**/

func main() {
	example.ExampleServiceClient_RegisterServiceInstance()
	example.ExampleServiceClient_Subscribe()
	time.Sleep(100*time.Second)
}
