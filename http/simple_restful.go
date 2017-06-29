package main

import (
	"fmt"
	"io"
	"net/http"
)

type myHandle int

/* 请求图片 */
func (handle myHandle) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
		case "/cat":
			io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">`)
		case "/dog":
			io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	}
}

func main() {
	var handle myHandle
	fmt.Println("Start Server")
	http.ListenAndServe(":8090", handle)
}
