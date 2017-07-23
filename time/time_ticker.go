package main

import (
	"fmt"
	"time"
)

// time.Ticker : 每隔一段时间, 都会向通道发送一个当时时间
// 类似于定时器: 类似函数还有: time.After()
func main() {
	duration, err := time.ParseDuration("10s")
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	//设置每隔10s 都会打印ticker
	ticker := time.NewTicker(duration)
	
	defer ticker.Stop()
	
	for {
		select {
			case <- ticker.C:
				fmt.Println("ticker: ", time.Now())
			case <- time.After(2 * time.Second):
				fmt.Println("After: ", time.Now())
		}
	}
}
