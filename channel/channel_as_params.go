/* channel 作为参数 */
package main

import (
	"fmt"
	"sync"
)

func doubletimes(a ...int) chan int {
	times := make(chan int)
	go func() {
		for _, v := range a {
			times <- v * 2
		}
		close(times)
	}()

	return times
}

func incr(c chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			out <- v + 1
		}
		close(out)
	}()
	return out
}

func output(ch ...chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(ch))

	/* 保证每个channel原子性 */
	/*
	for _, c := range ch {
		go func(c chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	*/

	//使用匿名函数
	outfunc := func(c chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	for _, c := range ch {
		go outfunc(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	c1 := doubletimes(2, 3, 5, 7, 11, 14, 16, 34)
	
	a1 := incr(c1)
	a2 := incr(c1) 		
	
	for n := range output(a1, a2) {
		fmt.Println(n)	
	}
}
