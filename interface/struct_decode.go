//JSON字符串转化对应的结构数据
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name	string
	Age		int
	Job		string
}

func main() {
	var	p1 Person
	str := strings.NewReader(`{"Name": "ZhangShan", "Age": 30, "Job": "Engneer"}`)

	json.NewDecoder(str).Decode(&p1)

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Println(p1.Job)
}
