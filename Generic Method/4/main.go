package main

import "fmt"

type myType interface {
	int64 | string
}

func bubleSort[T myType] (value []T) []T {
	for i := 0; i < len(value) - 1; i++ {
		currentFirstIndex := i
		leastIndex := i
		for j := i; j < len(value); j++ {
			if value[leastIndex] > value[j] {
				leastIndex = j
			}
		}
		value[currentFirstIndex], value[leastIndex] = value[leastIndex], value[currentFirstIndex]
	}

	return value
}

func main() {
	intList := []int64{5, 2, 4, 3, 1, 6, 0}
	stringList := []string{"agus", "budi", "eko", "dimas"}

	intListSorted := bubleSort(intList)
	stringListSorted := bubleSort(stringList)

	fmt.Println(intListSorted)
	fmt.Println(stringListSorted)
}