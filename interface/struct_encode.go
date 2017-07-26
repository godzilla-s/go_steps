package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name	string
	Age		int
	Job		string
}

func main() {
	p1 := Person{Name: "ZhangShan", Age: 30, Job: "Engneer"}

	json.NewEncoder(os.Stdout).Encode(p1)
}
