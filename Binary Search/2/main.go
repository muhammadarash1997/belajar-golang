// Go program to illustrate send
// and receive operation
package main

import "fmt"

func main() {
	daftar := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ukuran := len(daftar) - 1
	inputValue := 3
	max := ukuran
	var mid int
	min := 0
	var index int

	for {
		if max%2 == 0 {
			mid = ((max - min) / 2) + min

			if inputValue == daftar[mid] {
				index = mid
				break
			} else if inputValue > daftar[mid] {
				min = mid + 1
			} else if inputValue < daftar[mid] {
				max = mid - 1
			}
		} else {
			mid = ((max - 1 - min) / 2) + min

			if inputValue == daftar[mid] {
				index = mid
				break
			} else if inputValue > daftar[mid] {
				min = mid + 1
			} else if inputValue < daftar[mid] {
				max = mid - 1
			}
		}

	}

	fmt.Println(index)
}
