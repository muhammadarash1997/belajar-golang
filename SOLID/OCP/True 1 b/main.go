package main

import (
	"fmt"
)

type Operation interface {
	SetLeft(left float64)
	SetRight(right float64)
	GetLeft() float64
	GetRight() float64
	Calculate() float64
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

func (this *Addition) Calculate() float64 {
	return this.left + this.right
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

func (this *Subtraction) Calculate() float64 {
	return this.left - this.right
}

func main() {
	a := NewSubtraction()
	a.SetLeft(3)
	a.SetRight(5)

	b := NewAddition()
	b.SetLeft(3)
	b.SetRight(5)

	fmt.Println(a.Calculate())
	fmt.Println(b.Calculate())
}

// This is OCP because when we want to add new Operation like Multiply then we do not have to change Calculate method
// solidOCP_true1b : Reference Dicoding