//使用ServeMux

package main

import (
	"log"
	"io"
	"net/http"
)

func query(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "query operator action")
}

func update(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update operator action")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/query", query)
	mux.HandleFunc("/update", update)

	//注意第二参数
    err := http.ListenAndServe(":8001", mux)
    if err != nil {
        log.Fatal(err)
    }
}
