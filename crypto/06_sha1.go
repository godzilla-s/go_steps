package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	hv := sha1.New()
	hv.Write(data)
	fmt.Printf("SHA1: %x\n", hv.Sum(nil))
	fmt.Printf("OrgData: %s\n", data)
}
