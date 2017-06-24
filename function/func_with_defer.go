package main

import (
	"fmt"
)

//同一个函数内多个defer， 按照声明的先后顺序，从后往前执行
func call_test02() {
	defer func() {
		fmt.Println("call test02 function is close")
	}()

	defer func() {
		fmt.Println("going close test02 function")
	}()

	fmt.Println("finish call_test02")
}

func test01() {
	defer func() {
		fmt.Println("test 01 function is close")
	}()

	call_test02()

	fmt.Println("finish test01")
}

func main() {
	test01()
}
