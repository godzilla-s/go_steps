//简单的http服务

package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
)
/*
	http.ResponseWriter  返回信息
	http.Request		 客户端请求数据
*/
func PrintInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Printf("r.Form: %T, %v\n", r.Form, r.Form) //存储HTTP Get, Post参数
	fmt.Println("Path:", r.URL.Path)  // HTTP请求路径
	fmt.Println("scheme:", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //返回给客户端
	w.Write([]byte("很高兴见到你"))
}

func main() {
	//处理虚拟文件
	http.HandleFunc("/", PrintInfo)
	//创建监听
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


