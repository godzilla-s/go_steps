package main

import (
	"fmt"
)

func change(s []byte) {
	s = s[:0]
	for i:=0; i<100; i++ {
		if i % 2 == 0 {
            s = append(s, 'G')
        } else {
            s = append(s, 'K')
        }
	}
}

func main() {
	a := []byte{'a', 'b', 'c', 'd'}

	fmt.Printf("a: %s, len: %d, cap: %d\n", a, len(a), cap(a))

	a = a[:0]

	for i:=0; i<1000; i++ {
		if i % 2 == 0 {
			a = append(a, 'X')
		} else {
			a = append(a, 'O')
		}
	}

	fmt.Printf("a: %s, len:%d, cap: %d\n", a, len(a), cap(a))

	change(a)
	fmt.Println("after change:")
	fmt.Printf("a: %s, len:%d, cap: %d\n", a, len(a), cap(a))

	b := a[:0]
	for i:=0; i<100; i++ {
		if i % 2 == 0 {
			b = append(b, 'A')
		} else {
			b = append(b, 'B')
		}
	}

	fmt.Printf("b: %s, len: %d, cap: %d\n", b, len(b), cap(b))

	//append 内存变化
	fmt.Println("append 内存变化")
	c := make([]byte, 1)
	for i:=0; i<200000; i++ {
		c = append(c, 'V')
		if len(c) == cap(c) {
			fmt.Println(len(c))
		}
	}
}
