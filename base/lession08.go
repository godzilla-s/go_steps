//面向对象 

package main

import "fmt"

type Person struct {
	name string
	age	 int
}

type Student struct {
	Person
	college string
}

type Teacher struct {
	Person
	subject	string
	level	int
}

func (s Student) print() {
	fmt.Println(s.Person.name, s.Person.age, s.college)
}

func (t Teacher) print() {
	fmt.Println(t.Person.name, t.Person.age, t.subject, t.level)
}

func main() {
	stu := Student{Person{"James Brown", 25}, "Stanford"}
	stu.print()

	tch := Teacher{Person{"Donald Pane", 39}, "Physic", 3}
	tch.print()
}

