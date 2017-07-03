// 文件处理
package main

import (
	"fmt"
	"os"
)

func test(){
	os.Mkdir("file", 0777)
	os.MkdirAll("file/test/test1", 0777)
	fmt.Println("mkdir OK")

	err := os.Remove("file")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	
	os.RemoveAll("file")
}

func main() {
	test()
}
