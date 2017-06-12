package main


import (
	"fmt"
)


func main() {
	var emplyee = make(map[int]string)
	
	emplyee[1] = "Matt"
	emplyee[2] = "Josh"
	emplyee[3] = "王大锤"

	fmt.Println(emplyee)

	cars := map[string]float32{}
	cars["宝马"] = 55.6
	cars["劳斯莱斯"] = 120.5
	cars["奥迪"] = 48.6
	fmt.Println(cars)
}
