package main

import (
	"sync"
	"time"
	"fmt"
	_ "sync"
)

func Test() {
	fmt.Println("This is test")
}

func Run(i int) {
	Test()
	fmt.Println("Run is over: ", i)
}

var once = sync.Once{}
// Test 只运行一次
func Run2(i int) {
	once.Do(Test)

	fmt.Println("Run2 is over: ", i)
}

func main() {
	for i:=0; i<5; i++ {
		//go Run(i)
		go Run2(i)
	}

	time.Sleep(2 * time.Second)
}