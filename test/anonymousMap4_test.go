package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousMap4(t *testing.T) {

	type keys map[int]string // keys merupakan tipe alias dari map[int]string
	a := keys{1: "aa"}       // a merupakan objek baru bertipe keys
	b := keys{1: "bb"}       // b merupakan objek baru bertipe keys
	c := keys{1: "cc"}       // c merupakan objek baru bertipe keys

	result1 := a
	result2 := a[1]
	fmt.Println(result1) // map[1:aa]
	fmt.Println(result2) // aa
	assert.Equal(t, a, result1, "Result must be %v", result1)
	assert.Equal(t, a[1], result2, "Result must be %v", result2)

	result3 := b
	result4 := b[1]
	fmt.Println(result3) // map[1:aa]
	fmt.Println(result4) // aa
	assert.Equal(t, b, result3, "Result must be %v", result3)
	assert.Equal(t, b[1], result4, "Result must be %v", result4)

	result5 := c
	result6 := c[1]
	fmt.Println(result5) // map[1:aa]
	fmt.Println(result6) // aa
	assert.Equal(t, c, result5, "Result must be %v", result5)
	assert.Equal(t, c[1], result6, "Result must be %v", result6)
}
