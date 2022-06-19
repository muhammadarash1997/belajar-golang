// return an anonymous function from another function
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Returning anonymous function
func print() func(i, j string) string {
	ourFunc := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return ourFunc
}

func TestAnonymousFunction5(t *testing.T) {
	value := print() // variable value akan menampung sebuah fungsi atau menjadi fungsi
	result := value("Welcome ", "to ")

	fmt.Println(result)
	assert.Equal(t, "Welcome to GeeksforGeeks", result, fmt.Sprintf("Result must be %v", result))
}
