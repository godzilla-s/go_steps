// 函数
package main

import "fmt"

func test1(x, y int) (int, string) {
	n := x + y
	s := fmt.Sprintf("result is: %d", n)
	return n, s
}

func test() (int, int) {
	return 10, 15
}

func add(a, b int) (c int) {
	defer func() {
		c += 100
	}()  // 函数退出时执行

	c = a + b
	return 
}

// 多参数，类型确定
func test_sum(s string, n...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}

// 多参数, 类型未确定
func test_args(v ...interface{}) {
	fmt.Println(v[0], v[1], v[2])
}

// 匿名函数
func test_anonymous_func() {
	fn := func() { println("这是匿名函数") }
	fn()

	// 匿名函数数组
	fns := [](func(x int) int){
		func(x int) int { return x + 10 },
		func(x int) int { return x * 14 },
	}

	a := fns[0](25)
	fmt.Printf("func[0] 100: %d\n", a)

	//作为结构体成员
	d := struct {
		fn	func() string
	} {
		fn: func() string { return "This is struct element" },
	}

	println(d.fn())

	//channel 函数
	fc := make(chan func() string, 2)
	fc <- func() string { return "Channel Function" }
	println("channle fucntion: ", (<-fc)())
}

// 函数作为参数
func test_function_args(a, b int, sum func(int, int) int) int {
	return sum(a, b)	
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(test1(2,4))

	x, _ := test()
	fmt.Println(x)

	n := add(5, 6)
	fmt.Println("n is ", n)

	fmt.Println(test_sum("result is : %d", 2, 4, 6, 8 ,10))
	a := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(test_sum("sum of array: %d", a...))

	test_anonymous_func()
	
	test_args("Mi amo", 24234, "dingdong")
	
	//测试函数参数
	f := sum
	fmt.Println("测试函数参数: sum: ", test_function_args(3, 5, f))
}
