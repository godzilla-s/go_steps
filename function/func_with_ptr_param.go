package main

import "fmt"

type Person struct {
	id 		int
	name	string
	age		int
	height	float32
	weight	float32
}

/* 传入指针可以改变值 */
func changeValue(a int, s *string) {
    fmt.Println("in params: ", *s)
	*s = "changevalue"
	a = 20
}

func changeValue2(a *int, s *string){
	*s = "Changed"
	*a = 30
}

func PersonInit(p Person) {
	p.id = 100
	p.name = "HelloCrazy"
	p.age = 89
	p.height = 178.5
	p.weight = 75.5
}

func PersonInit2(p *Person) {
	p.id = 100
    p.name = "HelloCrazy"
    p.age = 89
    p.height = 178.5
    p.weight = 75.5
}

func main() {
	a := 12
	val := "RollWorld"
	changeValue(a, &val)
	fmt.Println("after value: ", val)

	changeValue2(&a, &val)
	fmt.Println("after change2: ", a, val)

	var p1 Person

	PersonInit(p1)
	fmt.Println(p1)
	PersonInit2(&p1)
	fmt.Println(p1)
}
