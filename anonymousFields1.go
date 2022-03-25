// In a Go structure, you are allowed to create anonymous fields.
// Anonymous fields are those fields which do not contain any name you
// just simply mention the type of the fields and Go will automatically
// use the type as the name of the field. You can create anonymous
// fields of the structure using the following syntax. Example below:

// type struct_name struct {
// 	int
//	bool
//	float64
// }

// In a structure, you are not allowed to create two or more fields
// of the same type as shown below:

// type student struct {
//	int
// 	int
// }

// You are allowed to combine the anonymous fields with the named fields
// as shown below:

// type student struct {
// 	name  int
// 	price int
// 	string
// }

// EXAMPLE
package main

import "fmt"

type student struct {
	int
	string
	float64
}

func main() {
	value := student{123, "Bud", 8900.23}

	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)
}
