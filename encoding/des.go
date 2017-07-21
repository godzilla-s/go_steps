package main

import (
	"fmt"
	"crypto/des"
)

func main() {
	cp, err := des.NewTripleDESCipher([]byte("9D37F2AF79A3B42F9C37F2AF79A3B42F"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cp)
}
