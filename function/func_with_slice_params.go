//数组参数

package main

import (
	"fmt"
)

func sum(a []int32) int32{
	
	var s int32
	s = 0

	for _, v := range a {
		s += int32(v)
	}

	return s
}

func sum2(v ...int) int {
	s := 0
	for _, v := range v {
		s += v
	}
	return s
}

func main() {
	a := sum([]int32{3, 4, 5, 6 ,7 ,8})
	fmt.Println("a: ", a)

	d := []int{6, 8, 9, 10, 23}
	b := sum2(d...)
	fmt.Println("b: ", b)
}
