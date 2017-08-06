// 锁与互斥

package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
	"runtime"
	"math/rand"
)

// 利用gotoutine简单测试
func test() {
	var opt uint64 = 0
	
	for i:=0; i<50; i++ {
		go func() {

			for {
				atomic.AddUint64(&opt, 1)  /* 使opt 自动加1 */
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(2 * time.Second)  /* 主线程 确定时间 */
	optRes := atomic.LoadUint64(&opt)
	fmt.Println("after add opt is ", optRes, " origin opt is ", opt)
}

// 互斥锁

func test_mutex() {
	state := make(map[int]int)
	var opt uint64 = 0

	var mutex = &sync.Mutex{}

	for i:=0; i<50; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&opt, 1)
				runtime.Gosched()
			}
		}()
	}

	for j:=0; j<100; j++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)

			mutex.Lock()
			state[key] = val
			mutex.Unlock()

			atomic.AddUint64(&opt, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)
	optVal := atomic.LoadUint64(&opt)

	fmt.Println("optVal is ", optVal)

	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}


func main() {
	//test()

	test_mutex()

}
