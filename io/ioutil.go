//使用ioutil包读取文本信息

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	//读取文件所有数据
	fmt.Println(string(buf))

	//直接使用ioutil包中的ReadFile读取文件
	buf2, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf2))
	
}
