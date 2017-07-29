//关键字new: 用来分配内存的内置函数
// make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配
package main

import (
	"fmt"
)

type Employee struct {
	name	string
	job		string
	salary	int
}

func main() {
	var e1 *Employee = new(Employee)	
	e1.name = "ZhangShan"
	e1.job = "Manager"
	e1.salary = 30000

	fmt.Printf("e1 type:%T, value: %v\n", e1, e1)
	fmt.Printf("e1 name: %s, e1 job: %s, e1 salary: %d\n", e1.name, e1.job, e1.salary)

	//new int类型
	var a *int = new(int)
	*a = 20
	fmt.Printf("a: type: %T, %v, %d\n", a, a, *a)

	var b *string = new(string)
	*b = "aaaaa"
	fmt.Printf("b type: %T, %v, %s\n", b, *b, *b)
}

