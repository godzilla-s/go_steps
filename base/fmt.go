// fmt 输入输出
package main

import (
	"fmt"
)

type Data struct {
 	vtype		string
	data		interface{}
	length		int
}

func main() {
	a := 100
	fmt.Printf("二进制: %b\n", a)
	fmt.Printf("八进制: %o\n", a)
	fmt.Printf("十进制: %d\n", a)
	fmt.Printf("十六进制: %x\n", a)
	fmt.Printf("%T:%v\n", a, a)
	fmt.Printf("Unicode: %U\n", a)
	
	b := 123.456789
	fmt.Printf("%T:%v\n", b, b)
	fmt.Printf("%f\n", b)
	fmt.Printf("%9f\n", b)
	fmt.Printf(".2f\n", b) 	//精度为2
	fmt.Printf("9.2f\n", b) //宽度为5， 精度为2
	
	c := "HellWorld"
	fmt.Printf("%T:%v\n", c, c)
	fmt.Printf("%#v\n", c)  //Go 表示
	
	d := Data {
		vtype: "string",
		data: "abcdefgh",
		length: 10,
	}
	fmt.Printf("%T:%v\n", d, d)
	fmt.Printf("%+v\n", d) 	//输出时会带字段名
	
	x := 20
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
