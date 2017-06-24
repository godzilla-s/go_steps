// 协程与线程
package main

import "fmt"
import "time"

func f(s string) {
	for i:=0; i<3; i++ {
		fmt.Println(s, ":", i)
	}
}

func test_goroutine() {
	go f("Running ..")  // 开启新的线程执行

	f("Fighting ..")  // 主线程执行

	go func(msg string) {
		fmt.Println(msg)
	} ("Inner call")

	time.Sleep(2 * time.Second)  // 睡眠2秒
}

//无缓冲
func sum(a []int, c chan int){  
	var total int
	for _, v := range a {
		total += v
	}	
	
	c <- total   //send to c
}

func test_channel() {
	a := []int{2, 3, 5 ,6 ,8, 12, 24, 5}
	
	c1 := make(chan int)  //无缓冲
	c2 := make(chan int)

	go sum(a[:4], c1)
	go sum(a[4:], c2)

	x := <-c1    //recv from c1, 默认情况下接受是阻塞的
	y := <-c2

	fmt.Println("x: ", x, "y: ", y)
}

// 有缓冲
func test_buffer_channel() {
	c := make(chan int, 3)
	c <- 2
	c <- 4
	c <- 12
	fmt.Println(<- c, <- c)
}

func fabn(n int, c chan int) {
	x, y := 1, 1
	for i:=0; i<n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)  // 如果不关闭就会死锁
}

func test_fabn() {
	c := make(chan int, 10)
	go fabn(cap(c), c)
	for a := range c {   // 不断去除channel里面数据, 知道channel被显示关闭
		fmt.Printf("%d ", a)
	}
	println()
}


//针对上面情况,使用select
func fabnachi(c, q chan int) {
	x, y := 1, 1

	for {	
		select {
			case c <- x:
				x, y = y, x+y
			case <- q:
				fmt.Println("Exit")
				return
			default:
				//fmt.Println("Block")   //阻塞，等待数据
		}
	}
}

func test_select() {
	c := make(chan int, 10)
	q := make(chan int)

	go func() {	
		for i:=0; i<10; i++ {
			fmt.Printf("%d ", <-c)	
		}
		q <- 0
	}()

	fabnachi(c, q)
}

//等到超时
func test_select_timeout() {
	c := make(chan int)
	o := make(chan bool)

	go func(){
		for {
			select {
				case v := <- c:
					println("recv: ", v)
				case <- time.After(5 * time.Second):
					println("timeout")
					o <- true
					break
			}
		}
	}()
	<- o
}

func main() {
	//test_goroutine()
	test_channel()
	test_buffer_channel()

	test_fabn()

	test_select()

	test_select_timeout()
}
