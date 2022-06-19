// Anonymous function as an argument
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Passing anonymous function
// as an argument
func cetak(paramAnony func(p, q string) string) string { // nama variable p q tidak mesti sama dengan yg ada di anonymous func yg akan dikirim
	return paramAnony("Geeks", "for")
}

func TestAnonymousFunction4(t *testing.T) {
	value := func(a, b string) string { // di sini variabel nya a b bukan p q, ini bisa karena yg dilihat adalah strukturnya, sama juga bisa, malah lebih rapi
		return a + b + "Geeks"
	}

	result := cetak(value)
	fmt.Println(result)

	assert.Equal(t, "GeeksforGeeks", result, "Result must be 'GeeksforGeeks'")
}
