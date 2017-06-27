//匿名函数

package main

import (
	"fmt"
)

func main() {
	f1 := func(a, b int) {
		fmt.Println(a, "+", b, "=", a + b)
	}

	f1(3, 4)

	//带返回值的匿名函数
	f2 := func(a, b int) int {
		fmt.Println("function multi")
		return a + b
	}

	v := f2(4, 7)
	fmt.Println(v)
}
