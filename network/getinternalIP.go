// 获取内网IP
package main 

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("fail to get interface addr:", err)
		return 
	}

	// IsLoopback() 回环地址: 即127.0.0.1
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet); 
		if ok {
			if ipnet.IP.IsLoopback() {
				continue 
			}
			if ipnet.IP.To4() != nil {
				fmt.Println("IP:", ipnet.IP.String())
			}
		}
	}
}