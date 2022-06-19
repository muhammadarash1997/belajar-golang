// DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD VAR

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person1 struct {
	name string
	age  int
}

func TestAnonymousStruct3(t *testing.T) {
	// DEKLARASI ANONYMOUS STRUCT DENGAN VAR
	var student1 struct {
		person1
		grade int
	}
	student1.person1 = person1{"wick", 21}
	student1.grade = 2

	result1 := student1.person1
	result2 := student1.grade
	fmt.Println(result1)
	fmt.Println(result2)
	assert.Equal(t, student1.person1, result1, "Result must be %v", student1.person1)
	assert.Equal(t, student1.grade, result2, "Result must be %v", student1.grade)

	// KELEMAHAN DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD VAR ADALAH TIDAK BISA LANGSUNG DIINISIALISASI SEPERTI DI BAWAH INI, DI BAWAH INI AKAN EROR
	// var student2 struct {
	// 	person1
	// 	grade int
	// }{
	// 	person1 = person1{"wick", 21}
	// 	grade = 2
	// }
}
