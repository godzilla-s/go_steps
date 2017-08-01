package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub		*Hub
	conn	*websocket.Conn
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
		_, buffer, err := cli.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("ReadMessage: ", err)
			}
		}

		cli.hub.broadcast <- buffer
	}
}

func (cli *Client) handleWrite() {
	defer func() {
		log.Println("client write exit")
		cli.conn.Close()
	}()

	for {
		select {
			case buffer := <-cli.send:
				err := cli.conn.WriteMessage(websocket.TextMessage, buffer)
				if err != nil {
					log.Println("client respinse message: ", err)
					break; 
				}
		}
	}
}

func serveChat(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade fail ", err)
		return 
	}
	
	cli := &Client{hub: hub, conn: conn, send: make(chan []byte, 512)}
	cli.hub.register <- cli

	go cli.handleRead()
	go cli.handleWrite()
}

