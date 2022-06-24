package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c", "d", "a", "c"}
	b := []int{1, 2, 3, 2, 1, 5}
	fmt.Println(Unique(a))
	fmt.Println(Unique(b))
}

func Unique[T int | string](list []T) []T {
	key := make(map[T]bool)
	new := []T{}
	for _, v := range list {
		if key[v] == false {
			key[v] = true
			new = append(new, v)
		}
	}

	return new
}
