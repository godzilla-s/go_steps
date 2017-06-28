package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8010")
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpConn, err := net.DailTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return 
	}

	tcpConn.Write([]byte("Hello"))
}
