//数组

package main

import (
	"fmt"
)

func main() {
	var a []int

	fmt.Printf("初始 length: %d, cap:%d\n", len(a), cap(a))
	a = append(a, 0)
	fmt.Printf("添加一个数据 length: %d, cap:%d\n", len(a), cap(a))
	for i:=0; i<10; i++ {
		a = append(a, i * 2 + 3);
		fmt.Printf("添加一个数据 length: %d, cap:%d, %v\n", len(a), cap(a), a)
	}
		
	for i:=1; i<8; i++ {
		a = append(a, i + 20)
		fmt.Printf("添加一个数据 length: %d, cap:%d, %v\n", len(a), cap(a), a)
	}

	fmt.Println("a :", a)
    fmt.Printf("length: %d, cap:%d\n", len(a), cap(a))
	//capacity值得变化: 初始化为 2^n 增长

	var b []string
	b = append(b, "Smile")
	fmt.Printf("字符串数组: length:%d, cap:%d, %v\n", len(b), cap(b), b)
	for i:=0; i<10; i++ {
		b = append(b, fmt.Sprintf("%d", (i+1)*100))
		fmt.Printf("字符串数组: length:%d, cap:%d, %v\n", len(b), cap(b), b)
	}
}
