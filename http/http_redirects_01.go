//重定向
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)

	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	const url = "/echo"
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusMovedPermanently)
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from server")
}

