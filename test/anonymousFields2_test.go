// TIPE PENULISAN 1 DAN 2 SAMA SAJA

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TIPE PENULISAN 1
// type student2 struct {
// 	int
// 	string
// 	float64
// }

// TIPE PENULISAN 2
type student2 struct {
	int     int
	string  string
	float64 float64
}

func TestAnonymousFields2(t *testing.T) {
	value := student2{123, "Bud", 8900.23}

	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)

	assert.Equal(t, 123, value.int, "Value int must be %v", value.int)
	assert.Equal(t, "Bud", value.string, "Value string must be %v", value.string)
	assert.Equal(t, 8900.23, value.float64, "Value float64 must be %v", value.float64)
}
