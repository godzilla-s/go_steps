package main

import (
	"fmt"
	"net/rpc"
)

// 客户端调用
func main() {
	client, err := rpc.DialHTTP("tcp", ":3003") 
    if err != nil { 
        fmt.Println("链接rpc服务器失败:", err) 
    } 
    var reply int 
    err = client.Call("Handle.GetInfo", 12, &reply) 
    if err != nil { 
        fmt.Println("调用远程服务失败", err) 
    } 
    fmt.Println("远程服务返回结果：", reply) 

}