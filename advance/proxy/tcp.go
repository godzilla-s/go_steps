package main

import (
	"fmt"
	"log"
	"net"
	"io"
)

type Proxy struct {
	proxyAddr	string
	backAddr	string
	lsner		net.Listener	
}

func (proxy *Proxy) Run () {
	lsn, err := net.Listen("tcp", proxy.proxyAddr)
	if err != nil {
		log.Fatal(err)
	}

	proxy.lsner = lsn

	fmt.Println("Server is On")
	for {
		cli, err := proxy.lsner.Accept()
		if err != nil {
			log.Println("Accept: ", err)
			continue
		}

		go proxy.serve(cli)
	}
}

func (proxy *Proxy) serve (cli net.Conn) {
	if proxy.backAddr == "" {
		log.Println("-b is required")
		return
	}

	//连接实际服务器
	server, err := net.Dial("tcp", proxy.backAddr)
	if err != nil {
		log.Fatal("Connect Real Server: ", err)
	}

	go func() {
		io.Copy(cli, server)
		server.Close()
		cli.Close()
	}()

	io.Copy(server, cli)
	cli.Close()
	server.Close()
}



