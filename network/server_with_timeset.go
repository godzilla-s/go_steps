//带定时的连接

package main

import (
	"net"
	"log"
	"time"
	"bufio"
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
		go handle_connect(c)
	}
}
