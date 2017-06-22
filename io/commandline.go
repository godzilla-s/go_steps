package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 128)

	fmt.Print(">:")
	fmt.Scanf("%s", &b)
	fmt.Println(string(b))
}
