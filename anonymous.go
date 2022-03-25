package main

import (
	"fmt"
)

func main() {
	// ANONYMOUS FUNCTION
	func() {
		fmt.Println("Hello")
	}()

	// FOR LOOP 1
	for i := 0; i < 5; i++ {
		// ANONYMOUS FUNCTION
		func() {
			fmt.Println(i)
		}()
	}

	// FOR LOOP 2 <-- sebaiknya digunakan untuk asynchronous, karena fmt.Println tidak langsung mengambil nilai i dari for tetapi dari anonymous function
	for i := 0; i < 5; i++ {
		// ANONYMOUS FUNCTION
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	fungsiSatu := func() {
		fmt.Println("Hello World")
	}
	fungsiSatu()

	var fungsiDua func() = func() {
		fmt.Println("Hay World")
	}
	fungsiDua()

	var fungsiTiga func(int, int) (int, float32)
	fungsiTiga = func(i1, i2 int) (int, float32) {
		return 100, 0.1
	}
	fungsiTiga(1, 2)

	var fungsiEmpat func(int, int) int
	fungsiEmpat = func(i1, i2 int) int {
		return 100
	}
	fungsiEmpat(1, 2)
}
