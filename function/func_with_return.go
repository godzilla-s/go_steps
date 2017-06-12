/* 函数返回 */

package main

import (
	"fmt"
)

func add(a, b int) int {
	return a+b
}

func strcat(sa, sb string) string {
	return fmt.Sprintf("%s-%s", sa, sb)
}

/* 多值返回 */
func calc(a, b int, opr string) (string, int) {
	if opr == "+" {
		return "sum ", a+b
	} else if opr == "-" {
		return "sub ", a-b
	} else if opr == "*" {
		return "multi ", a * b
	} else if opr == "/" {
		return "div ", a / b
	} else {
		return "nil", 0
	}
}

func main() {
	fmt.Println("3 + 5 =", add(3, 5))
	fmt.Println("strcat: ", strcat("James", "Maron"))

	oprator, val := calc(4, 6, "*")
	fmt.Println(oprator, val)
}
