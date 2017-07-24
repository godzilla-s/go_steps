package main

import (
	"fmt"
)

type Hub struct {
	clients		map[*Client]bool
	addClient	chan *Client
	broadcast	chan []byte
}

func NewHub() *Hub {
	h := Hub{
		clients: make(map[*Client]bool)
		addClient: make(chan *Client)
		broadcast: make(chan []byte)
	}
	return &h
}

func (h *Hub) Run() {
	for {
		select {
			case client := <- h.addClient:
				h.clients[client] = true
			case message := <- h.broadcast:
				for cli, v := range h.clients {
					cli.send <- message
				}
		}
	}
}
