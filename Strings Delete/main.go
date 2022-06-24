package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abac"

	s = strings.Replace(s, "a", "", 2)

	fmt.Println(s)
}
