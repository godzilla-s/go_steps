package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i:=0; i< 5; i++ {
			c1 <- i + 10
		}
	}()

	go func() {
		for i:=0; i<5; i++ {
			c2 <- i + 20
		}
	}()

	for i:=0; i<9; i++ {	
		select {
			case v1 := <-c1:
				fmt.Println(v1)
			case v2 := <-c2:
				fmt.Println(v2)
		}
	}

}
