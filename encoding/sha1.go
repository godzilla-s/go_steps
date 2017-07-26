package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Println(sha1.Sum(data))
	fmt.Println("New Sha1: ", sha1.New())
}
