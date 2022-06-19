package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Student struct {
	Student string // nama fields bisa sama dengan nama class, tetapi akan dikenal berbeda
	int     int
	string  string
	float64 float64
}

func TestAnonymousFields3(t *testing.T) {
	value := Student{"University", 123, "Bud", 8900.23}

	fmt.Println("Grade : ", value.Student)
	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)

	assert.Equal(t, "University", value.Student, "Value student must be %v", value.Student)
	assert.Equal(t, 123, value.int, "Value int must be %v", value.int)
	assert.Equal(t, "Bud", value.string, "Value string must be %v", value.string)
	assert.Equal(t, 8900.23, value.float64, "Value float64 must be %v", value.float64)
}
