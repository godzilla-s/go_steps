package main

import (
	"log"
	"net/http"
)

var hub = NewHub()

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Not Support method: ", r.Method)
		return 
	}

	http.ServeFile(w, r, "index.html")
}

func chatServer(w http.ResponseWriter, r *http.Request) {
	serveChat(hub, w, r)
}

func main() {
	hub.Run()

	http.HandleFunc("/", index)	
	http.HandleFunc("/chat", chatServer)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
