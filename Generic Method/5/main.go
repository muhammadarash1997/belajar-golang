package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "abcdefghhgfedecba"

	slice := strings.Split(s, "")
	slice = Unique(slice)
	counts := []int{}

	for _, v := range slice {
		counts = append(counts, strings.Count(s, string(v)))
	}

	uniqueList := Unique(counts)
	sort.Ints(uniqueList)
	if len(uniqueList) == 1 {
		fmt.Println("YES")
		return
	}
	if len(uniqueList) == 2 {
		if uniqueList[1]-uniqueList[0] == 1 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
	return
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
