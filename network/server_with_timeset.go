//带定时的连接

package main

import (
	"net"
	"fmt"
	"log"
	"time"
	"bufio"
	"io"
)

func handle_connect(c net.Conn) {
	//设置时限
	err := c.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("TIME OUT")
	}
	defer c.Close()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		println(line)
	}
}

func handle_connect2(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 200)
	for {
		//设置读取数据超时时限为10s
		err := c.SetReadDeadline(time.Now().Add(10 * time.Second))
    	if err != nil {
        	log.Println("设置读取超时失败")
    	}
		_, err = c.Read(buf)
		if err != nil{
			if err == io.EOF {
				fmt.Println("client close")
				return 
			}
			log.Println("time out", err)
			continue
		}
		fmt.Println("Read: ", string(buf))
	}
}

func main() {
	lsn, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	defer lsn.Close()

	for {
		c, err := lsn.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Println("New Client Connect")
		go handle_connect2(c)
	}
}
