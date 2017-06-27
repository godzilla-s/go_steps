//使用io包的函数读写

package main

import (
	"fmt"
	"bytes"
)

func main() {
	//使用bytes包
	buf := bytes.NewBuffer([]byte("I am Reading from Text"))
	
	readbuf := make([]byte, buf.Len())
	
	//bytes包继承了Read方法
	n, _ := buf.Read(readbuf)

	fmt.Printf("reader: %s, len: %d\n", readbuf, n)
	
}
