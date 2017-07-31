package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

func createCode(s, key string) string {
	sha := sha256.New
	fmt.Println("sha: ", sha)
	h := hmac.New(sha, []byte(key))
	fmt.Println("hmac: ", h)
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println("create code: ", createCode("zhuweijin", "GodBlessYou"))
}
