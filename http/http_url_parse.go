package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")	
	fmt.Printf("val: %s\n", v)
}

func main() {
	http.HandleFunc("/", parse)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
