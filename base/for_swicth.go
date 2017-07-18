package main

import "fmt"


func test_if() {
	a := 3
	
	if n:="ABC"; a > 0 {
		println("a > 0 ", n[0])
	} else if a == 0 {
		println("a == 0 ", n[1])
	} else {
		println("a < 0 ", n[2])
	}
}

func test_for() {
	// for 的用法
	a := "JUMP Loop"

	fmt.Printf("method for 1: ")
	//for i:=0; i<len(a); i++ {   len 多次调用
	for i, b := 0, len(a); i < b; i++ {
		fmt.Printf("%c ", a[i])
	}
	println()

	fmt.Printf("method for 2: ")
	n := len(a)
	for n > 0 {
		fmt.Printf("%c ", a[n-1])
		n--;
	}
	println()

	fmt.Printf("method for 3: ")
	for c := range a {
		fmt.Printf("%c ", a[c] )
	}
	println()

	fmt.Printf("method for 4: ")
	for _, d := range a {
		fmt.Printf("%c ", d)
	}
	println()

	// 数组
	e := []int{1, 2 ,3}
	for f, v := range e {
		if f == 0 {
			e[1], e[2] = 98, 99
			fmt.Println(e)
		}

		e[f] = v + 100
	}
	fmt.Println(e)

	h := []int{10, 20 ,30}
	for j, v := range h {
		if j == 0 {
			h = h[:3]		// 对slice修改不会影响range值
			h[2] = 110
		}
		println(j, v)
	} 
}

func test_switch() {
	i := []int{2, 3, 4}
	j := 3

	switch j {
		case i[0]:
			println("swicth is a")
		case 1, 3:
			println("switch is b")
			fallthrough	// 继续下一个 case
		case i[2]:
			println("swicth is c")
		default:
			println("swicth is d")
	}
}

// break 用于 for ,switch, select, continue只能用于 switch
func test_break_continue() {
	L1:
	for i:=0; i<5; i++ {
	L2:
		for j:=0; j<5; j++ {
			if j > 2 {
				continue L2
			}
			if i > 1 {
				break L1
			}
			print(i, ":", j, " ")
		}
		println()
	}
}

func main() {
	test_if()
	test_for()

	test_switch()

	test_break_continue()
}
