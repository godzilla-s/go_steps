package main

import (
	"fmt"
	"bytes"
)

func main() {
	var buffer1 bytes.Buffer

	for i:=0; i<5000; i++ {
		buffer1.WriteString("abc")
	}
	fmt.Println(buffer1)

	var buffer2 string
	for i:=0; i<5000; i++ {
		buffer2 = buffer2 + "abc"
	}
	fmt.Println(buffer2)
}
