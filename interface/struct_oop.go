// 接口
package main

import "fmt"
import "math"

//接口 多态
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	x, y float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.x * r.y
}

func (r rect) perim() float64 {
	return 2 * (r.x + r.y)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

//接口 作为参数
func exp(x interface{}) string {
	switch x:=x.(type) {
		case nil:
			return "NULL"
		case int, uint:
			return fmt.Sprintf("Int %d", x)
		case bool:
			if x { 
				return "ture"
			} else {
				return "false"
			}
		case string:
			return "String"
		default:
			return fmt.Sprintf("Unspected data type %T, %v", x, x)
	}
}

func test(){
	var a interface{}  // a 可以是任意类型
	fmt.Println(exp(a))

	a = 5
	fmt.Println(exp(a))
	a = true
	fmt.Println(exp(a))
	a = "I am Father"
	fmt.Println(exp(a))
	a = 4.56
	fmt.Println(exp(a))
}

func main() {
	r := rect{ x:2.3, y:4.5 }
	c := circle{ 4.8 }

	measure(r)
	measure(c)

	test()
}
