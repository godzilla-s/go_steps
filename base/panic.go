//异常的设计与捕获
package main

import (
	"fmt"
)

func test_add(a, b, c int) {
	if a + b < c || a + c < b || b + c < a {
		panic("impossible make triangle")
	}
}

func main() {
	//recover 可以补货异常，怎么处理交给程序员
	defer func() {
		fmt.Println("C")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("D")
		panic("end test main")
	}()

	go test_add(1, 1, 4)

	panic("after test_add exception")
}

