package main

import (
	"fmt"
)

type Person struct {
	name		string
	age			int
	height		float32
}

type Work struct {
	title		string
	dept		string
}

// map[key]value: key, value 可以为任意类型
func main() {
	person := map[Person]bool{}

	p1 := Person{"James", 26, 178.5}
	p2 := Person{"Walt", 33, 182.0}
	p3 := Person{"Mark", 31, 180.6}

	person[p1] = true
	person[p2] = true
	person[p3] = true

	fmt.Printf("person len:%d, %v\n", len(person), person)

	employ := make(map[Person]Work)
	employ[p1] = Work{"IT Engineer", "1001"}
	employ[p2] = Work{"Network Engineer", "1001"}
	employ[p3] = Work{"Accounting", "1003"}
	
	fmt.Println(employ)
}
