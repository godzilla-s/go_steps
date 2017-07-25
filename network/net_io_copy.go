package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	lsn1, err := net.Listen("tcp4", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer lsn1.Close()

	for {
		cli, err := lsn1.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("New Client")	

		//可以通过telnet host port来测试
		io.Copy(cli, cli)

		cli.Close()
	}
	
}


