// KOMBINASI SLICE & STRUCT

package main

import "fmt"

func main() {
	// KOMBINASI SLICE & STRUCT WITHOUT ANONYMOUS
	type person struct {
		name string
		age  int
	}
	var allStudents1 = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	for _, student := range allStudents1 {
		fmt.Println(student.name, "age is", student.age)
	}

	// KOMBINASI SLICE & STRUCT WITH ANONYMOUS
	var allStudents2 = []struct { // <-- []struct{} merupakan tipe data, ini sama seperti []int dari var s2 = []int{...}
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	for _, student := range allStudents2 {
		fmt.Println(student)
	}
}
