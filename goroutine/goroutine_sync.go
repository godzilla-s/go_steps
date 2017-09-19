package main

import "fmt"
import "sync"
import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
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
		wg.Done()
	}()
	wg.Wait()
	ticker.Stop()
}
