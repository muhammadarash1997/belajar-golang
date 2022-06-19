package main

import (
	"fmt"
)

type Operation interface {
	SetLeft(left float64)
	SetRight(right float64)
	GetLeft() float64
	GetRight() float64
}

type Addition struct {
	left   float64
	right  float64
	result float64
}

func NewAddition() Operation {
	return &Addition{}
}

func (this *Addition) SetLeft(left float64) {
	this.left = left
}

func (this *Addition) SetRight(right float64) {
	this.right = right
}

func (this *Addition) GetLeft() float64 {
	return this.left
}

func (this *Addition) GetRight() float64 {
	return this.right
}

type Subtraction struct {
	left   float64
	right  float64
	result float64
}

func NewSubtraction() Operation {
	return &Subtraction{}
}

func (this *Subtraction) SetLeft(left float64) {
	this.left = left
}

func (this *Subtraction) SetRight(right float64) {
	this.right = right
}

func (this *Subtraction) GetLeft() float64 {
	return this.left
}

func (this *Subtraction) GetRight() float64 {
	return this.right
}

type Calculator struct {
}

func (this *Calculator) Calculate(operation Operation) float64 {
	switch operation.(type) {
	case *Addition:
		return operation.GetLeft() + operation.GetRight()
	case *Subtraction:
		return operation.GetLeft() - operation.GetRight()
	default:
		return 0
	}
}

func main() {
	a := NewSubtraction()
	a.SetLeft(3)
	a.SetRight(5)

	b := NewAddition()
	b.SetLeft(3)
	b.SetRight(5)

	calculator := Calculator{}
	fmt.Println(calculator.Calculate(a))
	fmt.Println(calculator.Calculate(b))

}

// This is not OCP because when we want to add new Operation like Multiply then we also have to change Calculate method