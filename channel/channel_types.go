// 多类型的channel

package main

import (
	"fmt"	
)

type myFunc func(data string)

func main() {
	a := make(chan int)
	b := make(chan string)
	c := make(chan bool)
	d := make(chan struct{})
	e := make(chan myFunc) // 函数作为参数

	go func() {
		a <- 10
		b <- "Hello"
		c <- true
		e <- func(data string) {
			fmt.Println(data)
		}
		d <- struct{}{}	
	}()

	// chan struct{}  与 chan bool 相比,节省信道内存
	// chan struct{} 为 0
detected:
	for{
		select {
			case v := <- a:
				fmt.Println("chan int: ", v)
			case v := <- b:
				fmt.Println("chan string: ", v)
			case v := <- c:
				fmt.Println("chan bool: ", v)
			case <- d:
				fmt.Println("empty struct")
				break detected
			case v := <- e:
				v("HelloWorld") 
		}
	}

	fmt.Println("## end")
}