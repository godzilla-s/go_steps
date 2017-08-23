package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a 		int 
	b 		string 
	c		bool
}
func attr(d interface{}) {
	v_d := reflect.ValueOf(d)
	fmt.Println("ValueOf: ", v_d)

	t_d := reflect.TypeOf(d)
	fmt.Println("TypeOf: ", t_d)

	fmt.Println("Kind: ", t_d.Kind())
	fmt.Println("Name: ", t_d.Name())
	fmt.Println("String: ", t_d.String()) //返回类型的字符串表示
	fmt.Println("Size: ", t_d.Size()) //类型的值需要多少字节
	fmt.Println("Align: ", t_d.Align())
	//fmt.Println("Len: ", t_d.Len()) //返回array类型的长度, 否则panic
	if t_d.Kind() == reflect.Array || t_d.Kind() == reflect.Chan || t_d.Kind() == reflect.Map || 
		t_d.Kind() == reflect.Ptr || t_d.Kind() == reflect.Slice {
		fmt.Println("Elem: ", t_d.Elem()) //回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	}
	
	fmt.Println("Value Type: ", v_d.Type())
	
}

func main() {
	d := Data {
		a : 20,
		b : "ABsedfg",
		c : true,
	}

	attr(d)
}
