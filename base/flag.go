// flag包: 解析参数
package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	//os.Args 命令行所有的参数
	args := os.Args 

	fmt.Println(args)
	fmt.Println(args[1:]) //切片: 实际参数

	var buffer = flag.String("data", "type string", "This is New Message from Nokia")
	fmt.Println(buffer)
}
