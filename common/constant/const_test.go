package constant

import (
	"fmt"
	"testing"
)

/**
*
* @description : 
*
* @author : codezhang
*
* @create : 2019-01-08 18:16
**/

func TestChar(t *testing.T)  {
	fmt.Println("aaa")
	fmt.Println("a"+string(rune(2))+"a")
	fmt.Println("bbb")
	fmt.Println("b"+string(rune(1))+"b")
}