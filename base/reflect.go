//Go 类型反射
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "hello"
	
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	
	struct_reflect()
}


func struct_reflect() {
	x := struct {
		a 	int
		b	float32
		c	string
	} {
		a: 100,
		b: 2.345,
		c: "hello",
	}

	fmt.Println(reflect.TypeOf(x))
}
