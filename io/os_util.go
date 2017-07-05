//os 包基本函数的使用
package main

import (
	"fmt"
	"os"
)

func main() {
	hostName, _ := os.Hostname()
	fmt.Println("HostName: ", hostName)
	
	pageSize := os.Getpagesize()
	fmt.Println("PageSize: ", pageSize)

	//系统环境变量
	env := os.Environ()
	fmt.Println("OS Env: ", env)

	fmt.Printf("uid:%d, euid:%d, gid:%d egid:%d\n", os.Getuid(), os.Geteuid(), os.Getgid(), os.Getegid())

	//获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("dir:", dir)
}
