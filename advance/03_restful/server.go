package main

import (
	"log"
	"net/http"
)

func main() {
	routes := NewRouter()
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}

