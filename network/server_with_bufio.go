package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

//扫描数据流， 一行一行
func handleConnWithScanner(c net.Conn) {
	defer c.Close()
	
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		fmt.Println(">:", scanner.Text())
	}
}

//使用bufio Reader处理数据收发
func handleConnWithReader(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		//buf, _, err := reader.ReadLine()  不建议用
		buf, err := reader.ReadBytes('\n') 
		if err != nil {
			fmt.Println(err)
			break
		}
		
		fmt.Println("Read:", string(buf))
	} 
}

func main() {
	lsn, err := net.Listen("tcp", ":8010")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer lsn.Close()

	for {
		cli, err := lsn.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//go handleConnWithScanner(cli)
		go handleConnWithReader(cli)
	}
}
