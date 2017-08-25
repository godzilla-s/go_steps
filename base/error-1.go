package main

import (
	"fmt"
	"errors"
)

func main() {
	// 方法一
	var err error = errors.New("this is error message")
	fmt.Println(err)
	fmt.Println(err.Error())

	// 方法二
	err2 := fmt.Errorf("%s", "error message with fmt")
	fmt.Println(err2.Error())
}
