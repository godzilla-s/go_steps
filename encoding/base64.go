package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	//base64加密
	a := "This World Is changing, Blah Blah Blas"
	
	enCode := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	a64 := base64.NewEncoding(enCode).EncodeToString([]byte(a)) 
	
	fmt.Println(len(a), len(a64))
	fmt.Println(len(a64), a64)

	b := "Zhuweijin"
	c := "zhuWeiJin"

	b64 := base64.NewEncoding(enCode).EncodeToString([]byte(b))
	c64 := base64.NewEncoding(enCode).EncodeToString([]byte(c))
	fmt.Println(len(b64), b64)
	fmt.Println(len(c64), c64)

	a64_crypt := base64.StdEncoding.EncodeToString([]byte(a))
	fmt.Println(len(a64_crypt), a64_crypt)

	//base64解密
	a64_plain, err := base64.StdEncoding.DecodeString(a64_crypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(a64_plain), string(a64_plain))
}
