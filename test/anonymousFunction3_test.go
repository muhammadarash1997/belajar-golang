package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousFunction3(t *testing.T) {
	// Passing arguments in anonymous function
	result := func(ele string) string {
		fmt.Println(ele)
		return ele
	}("GeeksforGeeks")

	assert.Equal(t, "GeeksforGeeks", result, "Result must be 'GeeksforGeeks'")
}
