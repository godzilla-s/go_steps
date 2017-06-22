package main

import (
	"fmt"
)

func GetEvenNumber(a []int, callback func(int) bool) []int {
	var out []int
	for _, v := range a {
		if callback(v) {
			out = append(out, v)
		}
	}

	return out
}

func main() {
	va := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 13, 14}
	out := GetEvenNumber(va, func(n int) bool {
		if n % 2 == 0 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(out)
}
