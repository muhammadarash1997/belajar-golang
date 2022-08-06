package main

import "fmt"

func main() {
	var x *int

	z := 5

	x = &z

	fmt.Println(*x)
	fmt.Println(z)
}