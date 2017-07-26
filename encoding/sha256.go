package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	hs := sha256.New()

	hsVal := hs.Sum(nil)
	fmt.Printf("SHA256: %x\n", hsVal)

	hs2 := sha512.New()
	hsVal2 := hs2.Sum(nil)
	fmt.Printf("SHA512: %x\n", hsVal2)
}
