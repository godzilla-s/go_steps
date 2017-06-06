package main

import "fmt"
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

func main() {
	a := 12
	val := "RollWorld"
	changeValue(a, &val)
	fmt.Println("after value: ", val)

	changeValue2(&a, &val)
	fmt.Println("after change2: ", a, val)
}