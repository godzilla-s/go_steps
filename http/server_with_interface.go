//使用 interface 多态特性和ServeMux
package main

import (
	"log"
	"net/http"
	"io"
)

type queryWorker int
type updateWorker int

//ServeHTTP方法重写
func (wk queryWorker)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "query operator action")
}

func (wk updateWorker)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update operator action")
}

func main() {
	var query queryWorker	
	var update updateWorker

	mux := http.NewServeMux()
	mux.Handle("/query", query)
	mux.Handle("/update", update)

	//注意第二参数
	err := http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal(err)
	}
}
