//使用channel实现简单管道功能

package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i:=0; i<100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for v := range c1{
			c2 <- v
		}
		close(c2)
	}()

	for d := range c2 {
		fmt.Println(d)
	}
}
