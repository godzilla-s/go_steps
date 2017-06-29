package main

import (
	"net/http"
	"log"
)

func getpic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gtpl/dog.jpg") // 将文件读取出来返给客户端
	//注意 这个函数只能处理一次
	//http.ServeFile(w, r, "simple_httpserver.go") 
}

func getfile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "print_info.go")
}

func main() {
	http.HandleFunc("/", getpic)
	http.HandleFunc("/getfile", getfile)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
