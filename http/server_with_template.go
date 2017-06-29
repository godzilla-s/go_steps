package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

/* gtpl: go template*/

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("gtpl/index.gtpl")
	if err != nil {
		log.Fatal(err)
	}
	//返回网页
	tpl.Execute(w, tpl)
}

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
