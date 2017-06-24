package main

import (
	"fmt"
	"net"
	"log"
	_"io"
	"bufio"
)

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Println("New Connection")

	//读取客户端数据
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		ol := fmt.Sprintf("From Server: ", ln)
		fmt.Fprintln(c, ol)
	}
}

func main() {
	lsn, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}

	defer lsv.Close()

	for {
		cli, err := lsn.Accept()
		if err != nil {
			log.Fatal(err)
		}

		/*		
		//返回客户端，客户端可以通过telnet localhost port 连接
		io.Copy(cli, cli)
		cli.Close()
		*/

		go handleConn(cli)
	}
}


