package main

import (
	"fmt"
	"time"
)

// 函数作为参数channel
var (
	run = make(chan func(int64))
	quit = make(chan struct{})
)

func handle() {
	for {
		select {
			case fn := <-run:
				fn(time.Now().Unix())
				quit <- struct{}{}
				return 
		}
	}
}

func main() {
	fn := func(a int64) {
		fmt.Println("This is", a)
	}

	defer close(run)
	defer close(quit)

	go handle()
	run <- fn

	<- quit
}