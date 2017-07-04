package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Url Path:", r.URL.Path)
	//处理404
	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("Path:", r.URL.Path)
}

func main() {
	http.HandleFunc("/", parse)
	http.HandleFunc("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
