package main

import (
	"fmt"
	"net/url"
)

func main() {
	netstr := "http://www.baidu.com"

	u, err := url.Parse(netstr)
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("url parse ok")
	fmt.Println("Host: ", u.Host)
	fmt.Println("Scheme: ", u.Scheme)
}
