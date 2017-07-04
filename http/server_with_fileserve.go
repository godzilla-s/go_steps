//实现简单的文件服务
package main

import (
	"log"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir("/home/develop")))
	http.Handle("/c", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8001", nil)
    if err != nil {
        log.Fatal(err)
    }
}
