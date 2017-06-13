package main

import (
	"fmt"
)

func send(msg chan string) {
	msg <- "Hello this is my first channel"
	msg <- "china is a greate country"
	msg <- "with 5000 year history"
	msg <- "and 1.3 billion people"
	close(msg)
}

func main() {
	msg := make(chan string)
	go send(msg)

	//通过 range 来等待接受
	for s := range msg {
		fmt.Println(s)
	}
	
}