package main

import (
	"fmt"
	"hash/fnv"
	"os"
)

func main() {
	f, err := os.Open("hash_fnv.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	h := fnv.New64()

	fmt.Println(h)
}
