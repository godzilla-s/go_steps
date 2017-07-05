package main

import (
	"fmt"
)

func main() {
	var  buffer []string

	buffer = append(buffer, "0001")
	buffer = append(buffer, "0002")
	buffer = append(buffer, "0003")
	buffer = append(buffer, "0004")

	fmt.Println(buffer)

	buffer = buffer[:0] 	//clear buffer
	fmt.Println(buffer)
	
	buffer = append(buffer, "0005")
	fmt.Println(buffer)
}
