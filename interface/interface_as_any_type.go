package main

import (
	"fmt"
)

type Data struct {
	id   	int
	buf		string
}

func main() {
	vals := map[int]string {
		1: "Zhangshan",
		2: "Likui",
		3: "Wanermazi",
	}

	any := []interface{}{234, 456.7, vals, Data{23, "TestData"}, []int{2, 4, 6, 7}}

	fmt.Println(any)

	//map中value值未确定类型
	a := make(map[string]interface{})
	a["Action"] = "Beat"
	a["Hurt"] = true
	a["Rat"] = 123
	a["buffer"] = Data{10, "Unknow"}

	fmt.Println(a)
}
