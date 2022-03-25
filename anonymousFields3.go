package main

import "fmt"

type Student struct {
	Student string // nama fields bisa sama dengan nama class, tetapi akan dikenal berbeda
	int     int
	string  string
	float64 float64
}

func main() {
	value := Student{"University", 123, "Bud", 8900.23}

	fmt.Println("Grade : ", value.Student)
	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)
}
