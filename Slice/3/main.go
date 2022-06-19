package main

import (
	"fmt"
)

func main() {
	myArray := [6]int{10, 11, 12, 13, 14, 15}
	fmt.Println("myArray", myArray)

	slice1 := myArray[2:]
	fmt.Println("slice1", slice1)

	slice2 := slice1[2:4]
	fmt.Println("slice2", slice2)

	fmt.Println("myArray", myArray)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}
