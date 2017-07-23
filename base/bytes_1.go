package main

// bytes 包的使用

import (
	"fmt"
	"bytes"
)

func main() {
	a := []byte("Hello World")
	b := []byte("Hello World.")

	if bytes.Compare(a, b) == 0 {
		fmt.Println("a equal b")
	} else {
		fmt.Println("a unequal b")
	}

	s := []byte("Run the world, Run the Show, Run Run Run")

	fmt.Printf("count 'it' %d\n", bytes.Count(s, []byte("Run")))	

	s1 := []byte("   Smile Flower")
	fmt.Printf("trim space:%s\n", bytes.TrimSpace(s1))

	s2 := []byte("aaaYes to me aaaboaa")
	fmt.Printf("trim :%s\n", bytes.Trim(s2, "aob")) //去掉字符串中前后所有包含"aob"的字符

	// 同理 TrimLeft / TrimRight

	s3 := []byte("Dong Di Da Du Dr")
	fmt.Printf("Fileds String: %v\n", bytes.Fields(s3)) // 按空格切分字符串

	s4 := []byte("kiwi#apple#grape#pear")
	fmt.Printf("Replace string: %s\n", bytes.Replace(s4, []byte("#"), []byte(" "), -1))
}
