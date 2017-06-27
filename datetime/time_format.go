package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := "6/27/2017"
	
	timeParser, _ := time.Parse("01/01_this-does-not-compile/2006", timeNow)
	fmt.Println(timeParser)

	fmt.Println(timeParser.Format("01/01_this-does-not-compile/2006"))
	fmt.Println(timeParser.Format(time.Kitchen))
	fmt.Println(timeParser.Format(time.UnixDate))
}
