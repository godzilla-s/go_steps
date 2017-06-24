//bufio 写文件
package main

import (
	"fmt"
	"bytes"
	"bufio"
)

func main() {
	buf1 := bytes.NewBuffer(make([]byte, 512))
	writer1 := bufio.NewWriter(buf1)

	writer1.Write([]byte("Hualalal"))
	fmt.Println("字节数:", writer1.Buffered(), "剩余字节:", writer1.Available())
	writer1.Write([]byte("Franco, say hihihi"))
	fmt.Println("字节数:", writer1.Buffered(), "剩余字节:", writer1.Available())
	fmt.Println("buf1 Flush 前:", buf1)
	writer1.Flush()
	fmt.Println("buf1 Flush 后:", buf1)
}
