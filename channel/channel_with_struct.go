package main

import (
	"fmt"
	"time"
)

type Test struct {
	name	string
	id		int
	buf		string
}

func create(name string) chan Test {
	c := make(chan Test)
	go func() {
		for i:=0; i<10; i++{
			t := Test{name, i+1, fmt.Sprintf("Hello Welcome: %s", time.Now().String())}
			c <- t
		}
		close(c)
	}()

	return c
}

func run(c chan Test) {
	for v := range c {
		fmt.Println(v.name, v.id, v.buf)
	}
}

func main() {
	c := create("Jinggo")
	go run(c)
	go run(c)

	time.Sleep(1 * time.Second)
}
