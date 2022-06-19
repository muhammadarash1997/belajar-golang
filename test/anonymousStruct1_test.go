package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	name string
	age  int
}

func TestAnonymousStruct1(t *testing.T) {
	// Anonymous Struct
	// ini merupakan pembuatan struct sekaligus penciptaan objek
	var s1 = struct { // s1 merupakan sebuat object dari struct
		person
		grade int
	}{}
	// Inisialisasi
	s1.person = person{"wick", 21}
	s1.grade = 2

	result1 := s1.person.name
	result2 := s1.person.age
	result3 := s1.grade
	fmt.Println("name :", result1)
	fmt.Println("age :", result2)
	fmt.Println("grade :", result3)
	assert.Equal(t, s1.person.name, result1, "Result must be %v", result1)
	assert.Equal(t, s1.person.age, result2, "Result must be %v", result2)
	assert.Equal(t, s1.grade, result3, "Result must be %v", result3)

	// Anonymous Struct sekaligus Inisialisasi
	// ini merupakan pembuatan struct sekaligus penciptaan objek sekaligus inisialisasi
	var s2 = struct { // <-- struct{} merupakan tipe data, ini sama seperti []int dari var s2 = []int{...}
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	result4 := s2.person.name
	result5 := s2.person.age
	result6 := s2.grade
	fmt.Println("name :", result4)
	fmt.Println("age :", result5)
	fmt.Println("grade :", result6)
	assert.Equal(t, s2.person.name, result4, "Result must be %v", result4)
	assert.Equal(t, s2.person.age, result5, "Result must be %v", result5)
	assert.Equal(t, s2.grade, result6, "Result must be %v", result6)

}

// var s2 = struct { // <-- struct{} merupakan tipe data, ini sama seperti []int dari var s2 = []int{...}
// 	person
// 	grade int
// }{
// 	person: person{"wick", 21},
// 	grade:  2,
// }

// []int <-- artinya slice bertipe int
// []struct{...} <-- artinya slice bertipe struct{...}
