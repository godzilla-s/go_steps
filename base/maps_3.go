package main

import (
	"fmt"
)

func test_nil_map(v map[string]string) {
	fmt.Println("Actual map is nil")
	if _, ok := v["aaaa"]; ok {
		fmt.Println("Error Found")
	} else {
		fmt.Println("Found: ", ok)
	}
}

func complicated_map() {
	a := make(map[string]int)
	
	a["A001"] = 1
	a["B001"] = 2
	a["C001"] = 3
	
	b := make(map[string]map[string]int)
	b["BOX_A01"] = a
	
	c := make(map[string]map[string]map[string]int)
	c["HUGE_BOX_A01"] = b
	
	fmt.Println(c)
}

func main() {
	test_nil_map(nil)
}
