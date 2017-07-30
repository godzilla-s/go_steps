//字符串拼接
package main

import (
	"fmt"
	"bytes"
	"strings"
	"time"
)

var times = 50000

func main() {
	// 方法一： bytes.Buffer.WriteString
	t1 := time.Now()
	var buffer1 bytes.Buffer
	for i:=0; i<times; i++ {
		buffer1.WriteString("abc")
	}
	t2 := time.Now()
	fmt.Printf("byte.Buffer.WriteString time costs: %v\n", t2.Sub(t1))

	// 方法二: string + 
	var buffer2 string
	t3 := time.Now()
	for i:=0; i<times; i++ {
		buffer2 = buffer2 + "abc"
	}
	//fmt.Println(buffer2)
	t4 := time.Now()
	fmt.Printf("string + string time costs: %v\n", t4.Sub(t3))
	
	//方法三: strings.Join
	var buffer3 string
	t5 := time.Now()
	for i:=0; i<times; i++ {
		buffer3 = strings.Join([]string{buffer3, "abc"}, "")
	}
	t6 := time.Now()
	//fmt.Println(buffer3)
	fmt.Printf("string join time costs: %v\n", t6.Sub(t5))

	//方法四: fmt.Sprintf
	var buffer4 string
	t7 := time.Now()
	for i:=0; i<times; i++ {
		buffer4 = fmt.Sprintf("%s%s", buffer4, "abc")
	}
	//fmt.Println(buffer4)
	t8 := time.Now()
	fmt.Printf("fmt Sprintf time costs: %v\n", t8.Sub(t7))
}
