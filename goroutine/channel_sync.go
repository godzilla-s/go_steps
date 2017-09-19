package main

import "fmt"
import "time"

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		total := 0
		for t := range ticker.C {
			fmt.Println("Tick at:", t)
			total += 1
			if total > 10 {
				break
			}
		}
		done <- true
	}()
	<-done
	ticker.Stop()
}
