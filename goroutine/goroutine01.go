package main

import (
	"fmt"
	_ "time"
)

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar() {
	for i := 0; i < 35; i++ {
		fmt.Println("Bar: ", i)
	}
}

func main() {
	go foo()
	go bar()
	//主线程退出， 子线程同样退出，故没有输出
}
