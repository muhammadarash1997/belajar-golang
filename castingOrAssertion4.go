package main

import "fmt"

type myInt int

func main() {
	var i myInt = 4
	originalInt := int(i)

	fmt.Println(originalInt)
}