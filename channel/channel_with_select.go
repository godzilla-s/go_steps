package main

import (
	"fmt"
	_"time"
)

// golang select 类似于Linux C网络编程中 select, poll, epoll
// case 语句中只能处理 IO 操作


//不带缓存channel
func test_select1() {
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

	for i:=0; i<10; i++ {	
		select {
			case v1 := <-c1:
				fmt.Println("v1: ", v1)
			case v2 := <-c2:
				fmt.Println("v2: ", v2)
		}
	}
	close(c1)
	close(c2)
}

// deafault 当 channel 为空时默认处理
// 下面default 用来处理channel已满的情况
func test_select2() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	send := func(c chan int, incr int) {
		for i:=0; i<10; i++ {
			c <- i + incr
		}
		close(c)
	}

	go send(c1, 1)
	go send(c2, 11)

	for i:=0; i<30; i++ {
		select {
			case v1 := <-c1:
				fmt.Println("v1: ", v1)
			case v2 := <-c2:
				fmt.Println("v2: ", v2)
			default:
				fmt.Println("no handle")
		}
	}
}

func main() {
	test_select2()
}
