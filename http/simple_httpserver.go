package main

import (
	"fmt"
	"net/http"
	"io"
)

/* 下面用三种方法处理返回数据 */

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index tempalte running"))
}

func query(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Fruit, Airplane, Cars, Ships, Flowers")
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register OK")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/query", query)
	http.HandleFunc("/reg", register)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
