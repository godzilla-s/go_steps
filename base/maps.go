package main


import (
	"fmt"
)


func main() {
	/* 使用make定义 */
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

	/* map 迭代 */
	for key, val := range cars {
		fmt.Printf("车名: %s, 价格: %f\n", key, val)
	}

	/* 未确定值类型 */
	var xx  = map[string]interface{} {
		"name" : "马克保罗",
		"age" : 35,
		"salary" : 25000.00,
		"gender" : "male",
	}
	xx["living"] = "Shanghai"
	fmt.Println(xx)

}
