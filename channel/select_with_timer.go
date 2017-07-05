package main

import (
	"fmt"
	"time"
)

func main() {
	input := make(chan int, 1)
	quit := make(chan bool, 1)

	defer func() {
		fmt.Println("close channel")
		close(input)
		close(quit)
	}()

	n := 0

EXIT:
	for {
		select {
			case v := <-input:
				fmt.Println("v:", v)
				if v == 10 {
					quit <- true
				}
			case <-time.After(1 * time.Second):
				n++
				input<- n
			case <-quit:
				fmt.Println("Quit")
				break EXIT  //跳出循环
		}
	}

	fmt.Println("test finish!")
}

//跳出select循环
//1. break
//2. return


