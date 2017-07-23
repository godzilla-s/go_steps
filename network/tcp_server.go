//net 包 服务端的详细分析

package main

import (
	"fmt"
	"net"
	"time"
	"io"
)

func handleTCPConn(c *net.TCPConn) {
	defer c.Close()
	
	//设置读操作期限
	c.SetReadDeadline(time.Now().Add(10 * time.Second))

	/*
	//设置发送探测包属性, OS X 和 Linux 系统上，当一个连接空间了2个小时,
	//系统会以75秒间隔发送8个keepalive包，
	c.SetKeekAlive(true)
	//c.SetKeepAlivePeriod()
	*/

	data := make([]byte, 200)
	for {
		_, err := c.Read(data)
		if err != nil && err != io.EOF {
			//fmt.Println("err:", err)
			continue
		}
		c.Write([]byte("Hello"))
	}
}


func main() {
	// 
	ipaddr, err := net.ResolveTCPAddr("tcp4", ":8010")
	if err != nil {
		fmt.Println(err)
		return
	}

	//建立监听
	tcpLsn, err := net.ListenTCP("tcp4", ipaddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("网络:", tcpLsn.Addr().Network())

	
	for {
		tcpConn, err := tcpLsn.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Println("本地网络:", tcpConn.LocalAddr().String())
		fmt.Println("远程网络:", tcpConn.RemoteAddr().String())
	
		handleTCPConn(tcpConn)
	}
}
