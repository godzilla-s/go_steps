package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func() {
		time.Sleep(1*time.Second)
		c1 <- 2
		fmt.Println("send to channel")
	}()

	fmt.Println("Wait for recv")
	a := <-c1
	fmt.Println("a:", a)

	/* 如果下面c2为不带缓存的channel, 会出现什么? */
	//c2 := make(chan string)
	c2 := make(chan string, 3)

	c2 <- "NiHao"
	c2 <- "WoHenHao"
	c2 <- "Tianqi"

	close(c2)

	for v := range c2 {
		fmt.Println( v )
	}
}
