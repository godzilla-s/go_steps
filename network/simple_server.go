package main

import (
	"fmt"
	"net"
	"os"
	"io"
)

func handleClient(c net.Conn) {
	defer c.Close()

	io.WriteString(c, "Helle, Connection is on:")

	buffer := make([]byte, 256)

	for {	
		len, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Read err: ", err)
			return
		}
	
		if len == 0 {
			return 
		}
		fmt.Println("Read: ", string(buffer), len)
	}
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
		
		fmt.Println("New Client")

		go handleClient(cli)
	}
}
