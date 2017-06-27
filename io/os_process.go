//使用os包中的进程

package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	p, err := os.StartProcess("ll", {"-t"}, nil)
	if err != nil {
		log.Fatal(err)
	}
}
