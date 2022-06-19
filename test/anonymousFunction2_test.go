package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousFunction2(t *testing.T) {
	// Assigning an anonymous
	// function to a variable
	value := func() string {
		fmt.Println("Welcome! to GeeksforGeeks")
		return "Welcome! to GeeksforGeeks"
	}

	result := value()
	assert.Equal(t, "Welcome! to GeeksforGeeks", result, "Result must be 'Welcome! to GeeksforGeeks'")
}
