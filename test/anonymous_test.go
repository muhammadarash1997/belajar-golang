package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymous(t *testing.T) {
	// ANONYMOUS FUNCTION
	func() {
		fmt.Println("Hello")
	}()

	// FOR LOOP 1
	for i := 0; i < 5; i++ {
		// ANONYMOUS FUNCTION
		result1 := func() int {
			fmt.Println(i)
			return i
		}()
		assert.Equal(t, i, result1, fmt.Sprintf("Result must be %v", i))
	}

	// FOR LOOP 2 <-- sebaiknya digunakan untuk asynchronous, karena fmt.Println tidak langsung mengambil nilai i dari for tetapi dari anonymous function
	for i := 0; i < 5; i++ {
		// ANONYMOUS FUNCTION
		result2 := func(i int) int {
			fmt.Println(i)
			return i
		}(i)
		assert.Equal(t, i, result2, fmt.Sprintf("Result must be %v", i))
	}

	fungsiSatu := func() string {
		return "Hello World"
	}
	result3 := fungsiSatu()
	assert.Equal(t, "Hello World", result3, "Result must be %v", result3)
	

	var fungsiDua func() string = func() string {
		fmt.Println("Hay World")
		return "Hay World"
	}
	result4 := fungsiDua()
	assert.Equal(t, "Hay World", result4, "Result must be %v", result4)

	var fungsiTiga func(int, int) (int, float32)
	fungsiTiga = func(i1, i2 int) (int, float32) {
		return 100, 0.1
	}
	result5, result6 := fungsiTiga(1, 2)
	assert.Equal(t, 100, result5, "Result must be %v", result5)
	assert.Equal(t, float32(0.1), result6, "Result must be %v", result6)

	var fungsiEmpat func(int, int) int
	fungsiEmpat = func(i1, i2 int) int {
		return 100
	}
	result7 := fungsiEmpat(1, 2)
	assert.Equal(t, 100, result7, "Result must be %v", result7)
}
