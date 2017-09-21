// 简单工厂模式
package main

import (
	"fmt"
)

type Operation interface {
	calcResult() int
	SetValueA(int)
	SetValueB(int) 
}

type BaseElement struct {
	a int
	b int
}

func (n *BaseElement) SetValueA(_a int) {
	n.a = _a
}

func (n *BaseElement) SetValueB(_b int) {
	n.b = _b
}

type OperatorAdd struct {
	BaseElement
}

func (this *OperatorAdd) calcResult() int {
	return this.a + this.b
}

type OperatorSub struct {
	BaseElement
}

func (this *OperatorSub) calcResult() int {
	return this.a - this.b
}

type OperatorMul struct {
	BaseElement
}

func (this *OperatorMul) calcResult() int {
	return this.a * this.b
}

type OperationMod struct {
	BaseElement
}

func (this *OperationMod) calcResult() int {
	if this.b == 0 {
		fmt.Println("invalid")
	}
	return this.a / this.b
}

type OperatorFactory struct{
}

func (this *OperatorFactory) createOperator(opr string) (Operation) {
	var operator Operation
	switch opr {
		case "+":
			operator = new(OperatorAdd)
		case "-":
			operator = new(OperatorSub)
		case "*":
			operator = new(OperatorMul)
		case "/":
			operator = new(OperationMod)
		default:
			panic("error operator!")
	}
	return operator
}
func main() {
	var fn OperatorFactory
	oper := fn.createOperator("+")
	oper.SetValueA(20)
	oper.SetValueB(30)
	fmt.Println(oper.calcResult())

	oper = fn.createOperator("*")
	oper.SetValueA(23)
	oper.SetValueB(14)
	fmt.Println(oper.calcResult())
}