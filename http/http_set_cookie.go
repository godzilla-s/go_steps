package main

import (
	"fmt"
	"strconv"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/get", getCookie)
	http.HandleFunc("/add", add)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "username",
		Value: "zhuwj",
	})

	fmt.Fprintf(w, "set cookie ok!")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	v, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "fail to get cookie: ", err)
		return
	}

	fmt.Fprintln(w, "get cookie: ", v)
}

//使用cookie计数
func add(w http.ResponseWriter, r *http.Request) {
	v, err := r.Cookie("num")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name: "num",
			Value: strconv.Itoa(0),
		})
		fmt.Fprintln(w, "add: ", 0)
		return
	}

	n, err := strconv.Atoi(v.Value)
	if err != nil {
		fmt.Fprintln(w, "Not Integer")
		return
	}
	n++	
	v.Value = strconv.Itoa(n)
	
	http.SetCookie(w, v)

	fmt.Fprintln(w, "add: ", n)
}
