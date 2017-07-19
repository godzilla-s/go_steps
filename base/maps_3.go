package main

import (
	"fmt"
)

func test_nil_map(v map[string]string) {
	fmt.Println("Actual map is nil")
	if _, ok := v["aaaa"]; ok {
		fmt.Println("Error Found")
	} else {
		fmt.Println("Found: ", ok)
	}
}

func main() {
	test_nil_map(nil)
}
