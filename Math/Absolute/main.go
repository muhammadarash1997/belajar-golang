package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 3

	c1 := Abs(a - b)
	fmt.Println(c1)

	c2 := Abs(b - a)
	fmt.Println(c2)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
