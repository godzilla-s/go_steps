package main

import "fmt"

type Person struct {
	firstName	string
	lastName	string
	age			int
	gender		string
}

type Chinese struct {
	Person
	political	string
}

type American struct {
	Person
	lastName    string    /* 会被覆盖 */
	religion	string
}

func main() {
	p1 := Chinese{
		Person: Person{
			firstName: "Zhang",
			lastName: "Sanfeng",
			age:		 100,
			gender:  "male",
		},
		political: "masses",
	}

	p2 := American{
		Person: Person{
			firstName: "Jean",
			lastName: "William",
			age:	 65,
			gender:	 "female",
		},
		lastName: "Fatburg",
		religion: "Christian",
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age, p1.gender, p1.political)
	fmt.Println(p2.firstName, p2.lastName, p2.age, p2.gender, p2.religion)

	//new 创建的是个指针	
	p3 := new(Chinese)
	fmt.Printf("%T, %v\n", p3, p3)
	p3.Person.firstName = "Ma"
	p3.Person.lastName = "Ze"
	p3.political = "Advanced"

	fmt.Println(p3)
}
