package main

import (
	"fmt"
)

func main() {
	var a interface{} = "This is String"
	var b interface{} = 123.45
	var c interface{} = 200

	if str, ok := a.(string); ok {
		fmt.Println("string value:", str)
	} else {
		fmt.Println("not string")
	}

	if iv, ok := b.(float64); ok {
		fmt.Println("float32 value: ", iv)
	} else {
		fmt.Println("not float32 value")
	}

	//打印类型
	fmt.Printf("%T\n", c)
	fmt.Printf("add 100 : %d\n", c.(int) + 100)
	//不能用 fmt.Printf("add 100 : %d\n", a + 100)
}
