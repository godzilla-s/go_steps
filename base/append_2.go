package main

import (
	"fmt"
)

//append 只能处理 slice 数据
func main() {
	var s []byte

	s = append(s, "Hello"...)
	s = append(s, "World"...)
	s = append(s, "Smile"...)

	fmt.Printf("s: %s, len: %d, cap: %d\n", s, len(s), cap(s))
}
