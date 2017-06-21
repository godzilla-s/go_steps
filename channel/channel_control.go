package main

import (
	"fmt"
	"time"
)

type Control struct {
	num 	chan int
	data	chan string
	quit 	chan bool
}

func create() Control {
	return Control {
		make(chan int),
		make(chan string),
		make(chan bool),
	}
}

func pushmsg(ctl Control) {
	go func() {
		i := 1
		for {
			ctl.num <- i
			time.Sleep(1e8)
			i = i + 2
		}
	}()

	go func() {
		for {
			ctl.data <- time.Now().String()
			time.Sleep(2e8)
		}
	}()
}

func pullmsg(puller string, ctl Control) {
	go func() {
		for n := range ctl.num {
			fmt.Println(puller, "get number:", n)
		}
	}()

	go func() {
		for v := range ctl.data {
			fmt.Println(puller, "get data:", v)
		}
	}()
}

func timer(ctl Control) {
	time.Sleep(5 * time.Second)
	ctl.quit <- true
}

func main() {
	control := create()

	go pushmsg(control)
	go pullmsg("Polly", control)
	go pullmsg("Mashroom", control)
	go timer(control)

	<- control.quit
	fmt.Println("finish")
	close(control.num)
	close(control.data)
	close(control.quit)
}
