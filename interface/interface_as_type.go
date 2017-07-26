// interface{} 类型： 将实际的传入值或者定义的数据转化为interface{}类型
//而非 任何类型
package main

import (
	"fmt"
)


func main() {
	a := []interface{}{"你好", 200, 234.56}	
	fmt.Printf("a type:%T, %v\n", a, a)

	check_type(200)
	check_type(19.123)
	check_type("Hello")
}

func check_type(a interface{}) {
	
	switch a := a.(type) {
		case int:
			fmt.Println("type int: ", a)
		case bool:
			fmt.Println("type bool: ", a)
		case string:
			fmt.Println("type string: ", a)
		case []byte:
			fmt.Println("type []byte: ", a)
		case float32:
			fmt.Println("type float32: ", a)
		default:
			fmt.Println("undefine type: ", a)
	}
}
