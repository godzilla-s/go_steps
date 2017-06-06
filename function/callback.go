//回调函数
package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int 
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	visit([]int{1,2,3,4,5}, func(n int){
		fmt.Println(n)
	})

	xs := filter([]int{1,2,3,4,5}, func(n int) bool {
		return n>2
	})
	fmt.Println(xs)
}