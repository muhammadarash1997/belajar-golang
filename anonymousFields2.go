// TIPE PENULISAN 1 DAN 2 SAMA SAJA

package main

import "fmt"

// TIPE PENULISAN 1
// type student struct {
// 	int
// 	string
// 	float64
// }

// TIPE PENULISAN 2
type student struct {
	int     int
	string  string
	float64 float64
}

func main() {
	value := student{123, "Bud", 8900.23}

	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)
}
