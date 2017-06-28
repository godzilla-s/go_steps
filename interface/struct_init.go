package main

import (
	"fmt"
)

type Person struct {
	name	string
	age		int
	height	float32
	weight	float32
}

type Employee struct {
	Person
	title	string
	sal		float32
}

func EmployeeNew(name string, age int) *Employee {
	p := Person {
		name : name,
		age : age,
	}
	emp := &Employee{
		Person : p,
		sal : 20000,
	}
	return emp
}

func main() {
	emp := EmployeeNew("MaYun", 50)

	fmt.Printf("Type: %T, %v\n", emp, emp)

	emp.Person.height = 178.5
	emp.Person.weight = 70.0

	emp.title = "IT Engineer"

	fmt.Printf("Emp: %v\n", emp)
}
