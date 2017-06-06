/* 函数 */

package main

import (
    "fmt"
)

func test_add(a, b int) int {
	return a+b
}

func test_swap(a, b *int) {
	var temp int
	temp = *a;
	*a = *b
	*b = temp
}

func simple_swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("test a + b: ", test_add(2, 5))

	var a, b int
	a = 23
	b = 34
	fmt.Println("before swap: ", a, b)
	test_swap(&a, &b)
	fmt.Println("after swap: ", a, b)

	x, y := simple_swap(4, 7)
	fmt.Println("swap(4, 7): ", x, y)
}