// fmt 输入输出
package main

import (
	"fmt"
)

func main() {
	x := 20
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
