package main

import "fmt"

type myType interface {
	int64 | string
}

func display[T myType] (value T) {
	fmt.Println(value)
}

func main() {
	var myIntValue int64 = 12
	var myStringValue string = "Hello"
	display(myIntValue)
	display(myStringValue)
}