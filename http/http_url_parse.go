package main

import (
	"fmt"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Url Path:", r.URL.Path)
	//解析url参数
	v := r.FormValue("q")	
	fmt.Printf("val: %s\n", v)
	fmt.Println("Referer: ", r.Referer())
	// 客户端IP信息
	fmt.Println("Host addr: ", r.Host)
	// http 协议
	fmt.Println("Http Proco: ", r.Proto)
	// 请求的主体
	fmt.Println("Http Body: ", r.Body)
}

func main() {
	http.HandleFunc("/", parse)

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
