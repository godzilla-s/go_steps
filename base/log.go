package main

import (
	"log"
)

func main() {
	//进程退出
	log.Fatal("This Is Fatal log")

	//panic 线程错误退出
	log.Panic("This is Panic log")
}
