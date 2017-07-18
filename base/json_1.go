package main

import (
	"fmt"
	"os"
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

	//读取json文件数据
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(file)
	fmt.Println("decoder: ", decoder)
	
	err = decoder.Decode("Works")
	if err != nil {
		fmt.Println(err)
		return
	}
	
}


