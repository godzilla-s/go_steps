package main

import (
	"fmt"
	"net"
	"os"
)

func handleClient(c net.Conn) {
	defer c.Close()

	buffer := make([]byte, 256)
	len err := c.Read(buffer)
	if err != nil {
		fmt.Println("Read err: ", err)
		return
	}
	
	fmt.Println("Read: ", string(buffer))
}

func main() {
	lsn, err := net.Listen("tcp", ":8010")
	if err != nil {
		fmt.Println("Listen err: ", err)
		os.Exit(1)
	}

	defer lsn.Close()

	for {
		cli, err := lsn.Accept()
		if err != nil {
			fmt.Println("Accept err: ", err)
			os.Exit(1)
		}

		go handleClient(cli)
	}
}
