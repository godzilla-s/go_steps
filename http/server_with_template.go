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

func show(w http.ResponseWriter, r *http.Request) {
	tpl, err:= template.ParseFiles("gtpl/show.gtpl")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "gtpl/show.gtpl", "show")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}

