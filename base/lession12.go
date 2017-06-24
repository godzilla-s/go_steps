//socket 编程
package main

import (
	"fmt"
	_"os"
	"net"
)

func test() {
	service := "192.168.92.130:7200"

	addr := net.ParseIP(service)
	
	if addr != nil {
		fmt.Println("address is ", addr.String())
	} else {
		fmt.Println("Invalid address")
	}
}

func main() {
	test()
}


