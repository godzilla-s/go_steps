package main

import (
	"fmt"
)

type Data struct {
	id   	int
	buf		string
}

//根据switch判断数据类型
func valueType(v interface{}) {
	//fmt.Printf("type: %T\n", v)
	switch v:=v.(type) {		//这个只能用在switch结构内
		case string:
			fmt.Println("string:", v)
		case int:
			fmt.Println("int: ", v)
		case map[int]string:
			fmt.Println("map[int]string: ", v)
	}
}

func main() {
	vals := map[int]string {
		1: "Zhangshan",
		2: "Likui",
		3: "Wanermazi",
	}

	any := []interface{}{234, 456.7, vals, Data{23, "TestData"}, []int{2, 4, 6, 7}}

	fmt.Println(any)

	//map中value值为interface{}类型
	a := make(map[string]interface{})
	a["Action"] = "Beat"
	a["Hurt"] = true
	a["Rat"] = 123
	a["buffer"] = Data{10, "Unknow"}

	fmt.Println(a)

	valueType("Helle String")
	valueType(2000)
	valueType(vals)
}
