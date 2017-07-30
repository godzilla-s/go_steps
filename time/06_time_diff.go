package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(2 * time.Second)
	t2 := time.Now()

	fmt.Printf("time duration: %v\n", t2.Sub(t1))

	t3 := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Printf("time duration: %v\n", time.Since(t3))
}
