//简单的http服务

package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //返回给客户端
}

func main() {
	//处理虚拟文件
	http.HandleFunc("/", index)
	//创建监听
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


