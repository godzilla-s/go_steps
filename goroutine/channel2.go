package main

import (
	"fmt"
	_"time"
)

func worker(die chan bool) {
	fmt.Println("Begin Worker :")
	for {
		select {
			case <-die:
				fmt.Println("Done: this is worker:")
				die <- true
				return
		}
	}
}

func main() {
	die := make(chan bool)
	
	go worker(die)

	die <- true
	<- die
	fmt.Println("Worker goroutine has been terminated")
	//select {}  //会出现死锁
}