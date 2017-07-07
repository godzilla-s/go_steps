package main

import (
	"fmt"
	"control"
	"queue"
)

func main() {
	shutdown := make(chan bool)
	control := control.NewController()
	queue := queue.NewQueue(5, shutdown, control)
	
}
