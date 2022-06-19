package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousMap3(t *testing.T) {

	keys := map[int]string{
		1: "aa",
		2: "ab",
		3: "ac",
		4: "ba",
		5: "bb",
		6: "bc",
		7: "ca",
		8: "cb",
		9: "cc",
	}

	for i := 1; i <= len(keys); i++ {
		result := keys[i]
		fmt.Println(result)
		assert.Equal(t, keys[i], result, "Result must be %v", keys[i])
	}
}
