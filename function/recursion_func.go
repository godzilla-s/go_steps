// 递归函数
package main

import (
	"fmt"
)

func sum(a int) int {
	if a == 1 {
		return 1
	}

	return a + sum(a - 1)
}

func main() {
	fmt.Println(sum(1000000))
}
