package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	fmt.Println(string(data))

	
}
