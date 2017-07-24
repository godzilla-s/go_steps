package main

import (
	"fmt"
	"net"
)

type Client struct {
	hub		*Hub
	conn	*net.Conn
	send	chan []byte
}

func ClientNew(conn net.Conn) *Client {
	cli := Client {
		conn: conn,
		send: make(chan []byte),
	}
	
	return cli
}

func (cli *Client) handleRead() {
	defer func() {
		log.Println("client read exit")
		cli.hub.unregister <- cli
		cli.conn.Close()
	}()
	
	for {
		buffer, err := cli.conn.Read()
		if err != nil {
			log.Println("client read error")
			return
		}
		cli.hub.broadcast <- buffer
	}
}

func (cli *Client) handleWrite() {
	defer func() {
		log.Println("client write exit")
		cli.conn.Close()
	}

	for {
		select {
			case buffer := cli.send:
				err := cli.conn.Write(buffer)
				if err != nil {
					log.Println("write message : ", err)
					return
				}
		}
	}
}

func serveChat(hub *Hub, w http.ResponseWriter, r *http.Request) {
	cli := &Client{hub: hub, conn: , send: make(chan []byte, 256)}

	cli.hub.register <- cli
	
	go cli.handleRead()
	go cli.handleWrite()
}
