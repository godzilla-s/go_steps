package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

/* init，main() 函数是golang 保留的两个函数
   init 可以存在于任何包， main只能存在于main包中
   下边说一下两个函数的执行顺序：
   对同一个go文件的init()调用顺序是从上到下的
   对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数,对于
   对不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()
   如果package存在依赖，则先调用最早被依赖的package中的init()
   最后调用main函数
*/
func init() {
	// 开启多核运行
	fmt.Println("Number of CPU:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Print("Foo: ", i, "\n")
		time.Sleep(3 * time.Millisecond)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 25; i++ {
		fmt.Print("Bar: ", i, "\n")
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
