package main

import(
	"log"
	"net/http"
	_"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("/home/hdp/github"))
	
	http.Handle("/", fs)
	err := http.ListenAndServe(":8001", nil)
    if err != nil {
        log.Fatal(err)
    }
}
