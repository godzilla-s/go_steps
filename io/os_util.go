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
}
