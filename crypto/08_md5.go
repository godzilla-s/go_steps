package main

import (
	"fmt"
	"crypto/md5"
)

func MD5(data []byte) string {
	hv := md5.New()
	hv.Write(buf)
	return fmt.Sprintf("%x", hv.Sum(nil))
}

func main() {
	a := []byte("md5 string crypto")
	hv := md5.New()

	hv.Write(a)
	fmt.Printf("MD5 crypto: %x\n", hv.Sum(nil))
	b := []byte("123456789")
	fmt.Printf("MD5 crypto with buffer: %x\n", hv.Sum(b))
	
	fmt.Printf("MD 5: %s\n", MD5([]byte("123456789abcdefg")))
}

