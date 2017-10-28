package main

import (
	"fmt"

)

type fruit interface {
	Grow()
	Harvest()
	Sell()
}

type Apple struct {
}

func (a Apple) Grow() {
	fmt.Println("Grow apple")
}

func (a Apple) Harvest() {
	fmt.Println("Harvest apple")
}

func (a Apple) Sell() {
	fmt.Println("Sell apple")
}

type Banana struct {
}

func (b Banana) Grow() {
	fmt.Println("Grow banana")
}

func (b Banana) Harvest() {
	fmt.Println("Harvest banana")
}

func (b Banana) Sell() {
	fmt.Println("Sell banana")
}

type Farmer struct {
	fruit 
}

func NewFarmer(f fruit) *Farmer {
	return &Farmer{
		fruit: f,
	}
}

func (f *Farmer) Plant(x fruit) {
	f.fruit = x 
}

func main() {
	var a Apple 
	farmer := NewFarmer(a) 
	farmer.fruit.Grow()
	farmer.fruit.Harvest()
	farmer.fruit.Sell()

	var b Banana 
	farmer.Plant(b)
	farmer.fruit.Grow()
	farmer.fruit.Harvest()
	farmer.fruit.Sell()
}
