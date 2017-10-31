//内置函数copy
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 4}
	b := []int{10, 12, 13, 16}

	copy(b, a)
	fmt.Println("a:", a, "b:", b)

	var c []int

	copy(c, b)
	fmt.Println("copy b to c:", c)

	d := make([]int, 4)
	copy(d, a)
	fmt.Println(d)

	var xa [3]byte
	xb := []byte{97, 98, 99, 100}
	copy(xa[:], xb[:])  // 全部copy
	fmt.Println(xa)
}
