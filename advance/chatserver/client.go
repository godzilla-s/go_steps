package main

import (
	"log"
	"net"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub		*Hub
	conn	*net.Conn
	send	chan []byte
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
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

	cli.hub.register <- cli
	
	go cli.handleRead()
	go cli.handleWrite()
}
