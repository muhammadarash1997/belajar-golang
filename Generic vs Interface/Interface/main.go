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

func Unique(list interface{}) interface{} {
	key := make(map[interface{}]bool)
	new := []interface{}{}
	switch list.(type) {
	case []string:
		for _, v := range list.([]string) {
			if key[v] == false {
				key[v] = true
				new = append(new, v)
			}
		}
	case []int:
		for _, v := range list.([]int) {
			if key[v] == false {
				key[v] = true
				new = append(new, v)
			}
		}
	}

	return new
}