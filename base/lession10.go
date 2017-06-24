// 异常 

package main

import (
	"fmt"
)

func test(v ...interface{}) {
	fmt.Println(v[0], v[1], v[2])
}

func test_panic(a int){
	if a == 1 {
		fmt.Println("it is normal panic")
	}

	if a == 2 {
		panic("Undefine error")
	}

	if a == 3 {
		test_panic_recover()
	}
}

func test_panic_recover() {
	defer func() {
		fmt.Println("recover begin")
		if err := recover(); err != nil {  // 捕捉panic的错误, 不作exit处理
			fmt.Println(err)
		}
		fmt.Println("recover end")
	}()
	
	panic("test panic error")
}

func main() {
	test_panic(1)
	test_panic(3)
	test_panic(2)
}
