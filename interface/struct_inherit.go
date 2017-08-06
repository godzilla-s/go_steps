//继承

package main

import (
	"fmt"
)

type Person struct {
	name		string
	age			int
	sex			string
}

type Chinese struct {
	Person			//继承Person所有成员
	provice		string
	city		string
}

type American struct {
	*Person
	political	string
	state		string
	city		string
}


func main() {
	var c1 Chinese
	
	c1.name = "Zhang Shan"
	c1.age = 20
	c1.sex = "male"
	c1.provice = "Jiangxi"
	c1.city = "JiuJiang"

	fmt.Println(c1)

	var c2 American

	c2.Person = new(Person)

	c2.name = "Jack Baron"
	c2.age = 25
	c2.sex = "make"
	c2.political = "Public Party"
	c2.state = "California"
	
	fmt.Printf("c2: %v\n", c2)
}
