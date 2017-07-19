// 函数 方法 结构体

package main

import (
	"fmt"
	"math"
	//image/color"
)


type Point struct{
	x, y float64
}

func distance(b, e Point) float64 {
	return math.Hypot(b.x-e.x, b.y-e.y)
}

func (e Point) Distance (b Point) float64 {
	return math.Hypot(e.x-b.x, e.y-b.y)
}

func (e *Point) DistanceP (b Point) float64 {
	return math.Hypot(e.x-b.x, e.y-b.y)
}

func (a *Point) ScaleBy(times float64) {
	a.x = a.x * times
	a.y = a.y * times
}

func test() {
	a := Point{3.2, 4.0}
	b := Point{4.3, 1.4}

	fmt.Println("a and b distance: ", distance(a, b))  // 函数调用
	fmt.Println("a distance b: ", a.Distance(b))       // 方法调用

	c := &a
	fmt.Println("a distance b: ", c.DistanceP(b))		// 待指针的方法调用
	fmt.Println("a distance b: ", (&a).DistanceP(b))
	(&a).ScaleBy(2.0)
	fmt.Println("scale a to 2.0 ", a)
}

//继承
type ColoredPoint struct {
	Point			//继承上面
	color	int8
}

func (e ColoredPoint) Distance (b ColoredPoint) float64 {
	return e.Point.Distance(b.Point)
}

func (e *ColoredPoint) ScaleBy (times float64) {
	e.Point.ScaleBy(times)
}

func test2() {
	var p = ColoredPoint{Point{2.0, 3.0}, 1}
	var q = ColoredPoint{Point{5.0, 4.0}, 2}

	fmt.Println("test2: ", p.Distance(q))
	(&q).ScaleBy(2.3)
	fmt.Println("test2: scale :", q)
}

//方法值与表达式
func test3() {
	m := Point{3, 7}
	n := Point{9, 4}
	
	defDistanceFrom := m.Distance  // 方法值
	fmt.Println("distance from m to n: ", defDistanceFrom(n))
	var origin Point
	fmt.Println("origin distance: ", defDistanceFrom(origin))
/*
	defScaleMPoint := m.ScaleBy
	defScaleMPoint(3.1)
	fmt.Println("scale m: ", m)
*/
	//问题: defScaleMPoint之后，m 已经放大了, 而defDistanceFrom 还是初始的未Scaled的m 
	defDistance := Point.Distance
	fmt.Println("表达式: distance point m,n ", defDistance(n, m))
	fmt.Println("方法值: distance point m,n ", defDistanceFrom(n))

	defScale := (*Point).ScaleBy
	defScalePointM := m.ScaleBy
	defScale(&m, 3.1)
	fmt.Println("表达式: Scale point m ", m)
	defScalePointM(3.1)
	fmt.Println("方法值: Scale point m ❰", m, "❱")
}

/* 要点: 方法首字母大写，包外可访问；首字母非大写表示未导出，包外不可访问 */
func main() {
	test()
	test2()
	test3()
}

