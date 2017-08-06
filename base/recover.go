//panic 抛出异常; recover 捕获panic
package main

import (
	"fmt"	
)

func only_panic() {
	for i:=0; i<10; i++ {
		if i == 0 {
			panic("a is zero")
		}
		fmt.Printf("%d\t", i)
	}
	println()
	fmt.Println("end of only_panic")
}

func panic_with_recover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recover error: %v\n", err)
		}
	}()

	a := 10
	b := 0
	fmt.Println("a/b =", a/b)

	fmt.Println("end of panic_with_recover")
}

func main() {
	//only_panic()

	panic_with_recover()	
}
