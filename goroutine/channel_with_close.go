package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)  //自带3个缓存int
	c <- 13
	c <- 24
	c <- 30
	close(c)  //关闭
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	c <- 1  //会报错 send on closed channel
}
