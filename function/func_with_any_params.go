//任意参数, 类似于C中的void*

package main

import (
	"fmt"
)

func printmsg(v interface{}) {
	fmt.Println(v)
}

func printmultimsg(v ...interface{}) {
	fmt.Println(v)
}

func main() {
	printmsg(12)
	printmsg(34.56)
	printmsg("Smile ^o^")
	
	a := map[string]string {
		"Name": "DingDong",
		"Job": "Warning",
		"Weight": "12.3Kg",
	}
	printmsg(a)

	printmultimsg(23, 34.12, "font color", a)
}
