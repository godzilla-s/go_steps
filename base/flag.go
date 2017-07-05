// flag包: 解析参数
package main

import (
	"fmt"
	"flag"
)

func main() {
	a := flag.Int("num", 100, "your number")
	b := flag.String("name", "default", "your name")
	c := flag.Duration("du", 100, "your duration")

	flag.Parse()
	other := flag.Args()

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(other)
	fmt.Println(flag.Arg(2))
}
