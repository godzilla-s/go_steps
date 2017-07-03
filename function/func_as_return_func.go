// 函数作为返回值
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func undefine(a, b int) int {
	fmt.Println("Undefine function")
	return 0;
}

func calc(opr string) func(int, int) (int) {
	switch(opr) {
		case "+":
			return add
		case "-":
			return sub
		default:
			return undefine
	}
}

func main() {
	f1 := calc("+")
	fmt.Println(f1(23, 45))
	f2 := calc("-")
	fmt.Println(f2(43, 21))

}
