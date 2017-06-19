package main

import (
	"fmt"
	"time"
)

func test1(c chan int) {
	for i:=0; ; i++ {
		c <- i + 10
	}
}

func test2(c chan int) {
	for i:=0; ; i++ {
		c <- i + 2
	}
}

func scheule(c1, c2 chan int) {
	for {
		select {
			case v := <- c1:
				fmt.Println("Accept test1 chan: ", v)
			case v := <- c2:
				fmt.Println("Accept test2 chan: ", v)
		}
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go test1(c1)
	go test2(c2)

	go scheule(c1, c2)

	time.Sleep(1e5)	
}
