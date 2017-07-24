package main

import (
	"log"
)

type Hub struct {
	clients		map[*Client]bool
	register	chan *Client
	unregister	chan *Client
	broadcast	chan []byte
}

func NewHub() *Hub {
	h := Hub{
		clients: make(map[*Client]bool),
		register: make(chan *Client),
		unregister: make(chan *Client),
		broadcast: make(chan []byte),
	}
	return &h
}

func (h *Hub) Run() {
	for {
		select {
			case client := <- h.register:
				log.Println("Client login")
				h.clients[client] = true
			case client := <- h.unregister:
				log.Println("Client logout")
				delete(h.clients, client)
				close(client.send)
			case message := <- h.broadcast:
				for cli := range h.clients {
					select {
						case cli.send <- message:
						default:
							delete(h.client, cli)
							close(cli.send) 
					}
				}
		}
	}
}
