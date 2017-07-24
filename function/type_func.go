package main

import (
	"fmt"
)

type MyFunc func (a, b int)(string) 

func (f MyFunc) Add (a, b int) (string) {
	return fmt.Sprintf("calculate add %d and %d is %d", a, b, a+b)
}

func (f MyFunc) Multi (a, b int) (string) {
	return fmt.Sprintf("calculate multi %d and %d is %d", a, b, a * b)
}

func main() {
	var f1 MyFunc
	fmt.Println(f1.Add(3, 4))
	fmt.Println(f1.Multi(5, 67))
}
