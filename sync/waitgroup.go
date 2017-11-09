package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup
// 线程组
// Wait() 等待所有现场执行完成 
// Add(1) 新增一个线程
// Done() == Add(-1)完成执行一个线程 

func main() {
	var counter int 
	for i:=0; i<10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Hello: ", n)
			time.Sleep(3e9)
			wg.Done()
			// 或者: wg.Add(-1)
		}(i)  // 传入参数
	}
	fmt.Println(counter)
	wg.Wait()
}