package main

import (
	"fmt"
	"net"
)

type Client struct {
	conn	net.Conn
	send	chan []byte
}

func ClientNew(conn net.Conn) *Client {
	cli := Client {
		conn: conn,
		send: make(chan []byte),
	}
	
	return cli
}
