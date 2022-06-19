// KOMBINASI SLICE & STRUCT

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousStruct2(t *testing.T) {
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
		result1 := student.name
		result2 := student.age
		fmt.Println(result1, "age is", result2)
		assert.Equal(t, student.name, result1, "Result must be %v", student.name)
		assert.Equal(t, student.age, result2, "Result must be %v", student.age)
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
		result := student
		fmt.Println(result)
		assert.Equal(t, student, result, "Result must be %v", student)
	}
}
