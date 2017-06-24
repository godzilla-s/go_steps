package main

// type: slice, map, array 等结构类型

import "fmt"

//使用type重新定义变量名，可用于全栈
type bigint int64

func test(){
	var a bigint = 120	
	fmt.Println("a: ", a)

	b := 230
	var c bigint = bigint(b)  //不能直接 = a, bigint不是 int64 别名, 除了数据结构一样外，其他没有任何关系
	var d int64 = int64(c)
	println(b, c, d)

	
}

func array_test() {
	var e = []int{2, 3, 5}
	fmt.Println(e)
	fmt.Printf("e: %T\n", e)
}

func main() {
	test()
	array_test()
}
