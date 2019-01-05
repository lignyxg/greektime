package main

import (
	"bytes"
	"fmt"
)

//穷举a~e5个字母组成的4为密码
//var pwds = make([][]byte, 0)
var elems = []byte{'a','b','c','d','e'}
var secret = []byte{'d','e','a','a'}

func password(result []byte) {
	if len(result) == 4 {
		//pwds = append(pwds, result)
		if  bytes.Equal(secret, result) {
			fmt.Printf("find pwd: %s\n",result)
		}
		return
	}

	for i := 0; i < len(elems); i++ {
		resNew := make([]byte, len(result))
		copy(resNew, result)
		resNew = append(resNew, elems[i])
		password(resNew)
	}
}

func main() {
	res := make([]byte, 0)
	password(res)
}