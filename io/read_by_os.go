//使用os包读写文件
package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Open file error: ", err)
		os.Exit(1)
	}

	defer f.Close()

	buffer := make([]byte, 512)
	for {
		len, err := f.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("read error: ", err)
			os.Exit(1)
		}
		if len == 0 {
			fmt.Println("finish read")
			break
		}
	}

	fmt.Println(string(buffer))
}

