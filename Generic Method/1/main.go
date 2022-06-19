package main

import "fmt"

func display[T any] (value T) {
	fmt.Println(value)
}

func main() {
	var myIntValue int64 = 12
	var myStringValue string = "Hello"
	display(myIntValue)
	display(myStringValue)
}