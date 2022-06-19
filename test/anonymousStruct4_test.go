package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousStruct4(t *testing.T) {
	element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}

	result1 := element.name
	result2 := element.branch
	result3 := element.language
	result4 := element.Particles
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	assert.Equal(t, element.name, result1, "Result must be %v", element.name)
	assert.Equal(t, element.branch, result2, "Result must be %v", element.branch)
	assert.Equal(t, element.language, result3, "Result must be %v", element.language)
	assert.Equal(t, element.Particles, result4, "Result must be %v", element.Particles)
}
