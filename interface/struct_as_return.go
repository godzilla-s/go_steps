//struct 作为返回值

package main

import (
	"fmt"
)

type Person struct {
	name	string
	age		int
	sex		string
}

type Student struct {
	Person
	grade	string
}

type Teacher struct {
	Person
	subject		string
}

func NewStudent() Student {
	return Student{
		Person: Person{
			name: "Tomas",
			age: 23,
			sex: "male",
		},
		grade: "College Student",
	}
}

func main() {
	st := NewStudent()
	fmt.Println(st)
}
