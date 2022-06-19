package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {
	// pass a slice in variadic function
	element := []string{"geeks", "FOR", "geeks"}
	fmt.Println(joinstr(element...))
}
