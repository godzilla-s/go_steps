package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func() {
		time.Sleep(2*time.Second)
		c1 <- 2
		fmt.Println("send to channel")
	}()

	fmt.Println("Wait for recv")
	a := <-c1
	fmt.Println("a:", a)
}
