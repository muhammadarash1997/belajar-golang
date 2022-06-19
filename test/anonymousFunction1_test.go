package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousFunction1(t *testing.T) {
	// Anonymous function
	result := func() string {
		fmt.Println("Welcome! to GeeksforGeeks")
		return "Welcome! to GeeksforGeeks"
	}()

	assert.Equal(t, "Welcome! to GeeksforGeeks", result, "Result must be 'Welcome! to GeeksforGeeks'")
}
