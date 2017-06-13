package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

var mutex sync.Mutex // 锁

func main() {
	wg.Add(2)
	go incrementor3("Foo:")
	go incrementor3("Bar:")
	wg.Wait()

	fmt.Println("Final Counter: ", counter)
}

//不加锁， 竞态累加
func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Countor:", counter)
	}
	wg.Done()
}

//加锁，通过锁来保证变量原子性
func incrementor2(s string) {
	for i:=0; i<20; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Countor:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

//直接使用原子变量
func incrementor3(s string) {
	for i:=0; i<20; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)  //原子自增1
		fmt.Println(s, i, "Countor:", counter)
	}
	wg.Done()
}