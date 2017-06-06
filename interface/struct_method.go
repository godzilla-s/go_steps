package main

import "fmt"

type Person struct{
	name	string
	nation	string
	age		int
}

type Teacher struct {
	Person
	subject		string
}

func (p Person) ShowInfo() {
	fmt.Println("Name:", p.name, "Nation:", p.nation, "Age:", p.age)
}

func (t Teacher) ShowInfo() {
	fmt.Println("Name", t.name, "Nation:", t.nation, "Age:", t.age, "Subject:", t.subject);
}

func main() {
	p1 := Person{"Jason", "USA", 23}
	p2 := Teacher{Person{"Baron", "France", 24}, "Math"}

	p1.ShowInfo()
	p2.ShowInfo()
	p2.Person.ShowInfo()
}