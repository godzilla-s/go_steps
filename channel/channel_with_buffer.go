/* 带缓存的channel */
package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int, 5)

	go func() {
		for i:=0; i<20; i++ {
			c1 <- i
		}
		fmt.Println("send channel finish")
		close(c1)
	}()

	for i:=0; i<20; i++ {
		fmt.Println(<-c1)
	}
}
