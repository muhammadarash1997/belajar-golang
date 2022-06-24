package main

import "fmt"

func main() {
	name := "joko"
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		name := "joki"	// this variable can't be accessed from outside of the increment
		counter++
		fmt.Println(name)
	}

	increment()

	fmt.Println(counter, name)
}
