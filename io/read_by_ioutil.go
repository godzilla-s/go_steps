package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/* 直接读取 */
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	fmt.Println(string(data))

	/* 先打开，在读取 */
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)	
	if err != nil {
		fmt.Println("readall ", err)
		os.Exit(1)
	}
	fmt.Println("Read with os.ReadAll:\n", string(buf))
}
