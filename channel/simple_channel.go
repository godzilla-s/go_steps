package main

import (
	"fmt"
	"time"
)

func only_write(ch chan int) {
	for i:=0; i<10; i++ {
		ch <- i
	}
}

func only_read(ch <- chan int) {
	for v := range ch {
		fmt.Println("Read: ", v)
	}
}

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

	/*
	//单向读写的channel
	c0 := make(chan int)
	c3 := <-chan int(c0)  	//channel c0的单向读
	c4 := chan<- int(c0)	//channel c0的单向写
	*/

	//channel 机制是先进先出
	c5 := make(chan interface{}, 3)  //任意类型

	c5 <- "HelloWorld"

	fmt.Println(<-c5)

	//问题: 造成 dead lock：
	//如果是无缓冲的channel，往里面赋值时, 必须被读取出来
	//如下面情况：c6的读写是同步的， 先写在读，5被写进去，会阻塞
	/*
	c6 := make(chan int)
	c6 <- 5
	fmt.Println(<-c6)
	*/

	//上面例子建议使用缓冲，这样不会造成dead lock
	c6 :=  make(chan int, 1)
	c6 <- 5
	fmt.Println(<-c6)

}
