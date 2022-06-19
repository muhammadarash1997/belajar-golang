package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	// CARA INISIALISASI 1 DAN CARA INISIALISASI 2 AKAN MENGHASILKAN OUTPUT YANG BERBEDA

	// CARA INISIALISASI 1
	// var p1 = person{name: "wick", age: 21}
	// var s1 = student{person: p1, grade: 2}

	// CARA INISIALISASI 2
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)
}
