//实现简单的文件服务
package main

import (
	"log"
	"net/http"
)

func main() {
	//
	http.Handle("/", http.FileServer(http.Dir("/home/hdp")))

	err := http.ListenAndServe(":8001", nil)
    if err != nil {
        log.Fatal(err)
    }
}
