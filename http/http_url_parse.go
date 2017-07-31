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

func parseURL(w http.ResponseWriter, r *http.Request) {
	urlinfo := r.URL
	
	fmt.Println("Scheme: ", urlinfo.Scheme)
	fmt.Println("Opaque: ", urlinfo.Opaque)	//编码后的不透明数据
	fmt.Println("Host: ", urlinfo.Host)
	fmt.Println("RawQuery: ", urlinfo.RawQuery) //url查询数据 ?后面
	fmt.Println("Fragment: ", urlinfo.Fragment)	//#后面

	//http://192.168.92.130:8001/url?name=123&pass=123
}

func main() {
	http.HandleFunc("/", parse)
	http.HandleFunc("/url", parseURL)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
