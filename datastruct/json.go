package main

import (
	"fmt"
	"encoding/json"
)

/* 先定义一个科目对象 */
type Subject struct {
	name	string
	score	float32
}

/* 定义一个学生信息对象 */
type Student struct {
	name 		string
	subjects	[]*Subject
}

var bson = `{"name":"Taren", "subjects":[{"Math":90, "Physic":92}]}`

func main() {
	sj1 := &Subject{"Math", 100}
	sj2 := &Subject{"Chinese", 100}
	
	stu := Student{"张三丰", []*Subject{sj1, sj2}}

	js, _ := json.Marshal(stu)
	fmt.Println("Json:", js)

}


