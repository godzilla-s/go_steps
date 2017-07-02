//interface: 1. 类型的集合 2. 函数的集合(可用于多态)

package main

import (
	"fmt"
)


func main() {
	a := []interface{}{"你好", 200, 234.56}	
	fmt.Printf("a type:%T, %v\n", a, a)
}
