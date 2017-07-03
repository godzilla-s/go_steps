//版本 v1.7

//使用：
//不要将 Contexts 放入结构体，
//context 应该作为第一个参数传入，命名为 ctx
//不要传入 nil 的 Context
package main

import (
	"fmt"
	"time"
	"context"
)

func test_DeadLine() {
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
		case <- time.After(2 * time.Second):
			fmt.Println("over slept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
	}
}

func test_Timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancel()

	select {
		case <- time.After(1 * time.Second):
			fmt.Println("over slept")
		case <- ctx.Done():
			fmt.Println("time out")
	}
}

func main() {
	test_DeadLine()

	test_Timeout() 
}
