// Go program to illustrate send
// and receive operation
package main

import "fmt"

func main() {
	daftar := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ukuran := len(daftar) - 1
	inputValue := 10
	max := ukuran
	var mid int
	min := 0
	var index int

	for {
		yesyes := checkParitas(&index, &max, &min, &mid, &inputValue, &daftar)
		if yesyes == 1 {
			break
		}
	}

	fmt.Println(index)
}

func checkParitas(index *int, max *int, min *int, mid *int, inputValue *int, daftar *[10]int) int {
	if *max%2 == 0 {
		*mid = ((*max - *min) / 2) + *min
		yes1 := searchIndex(index, max, min, mid, inputValue, daftar)
		if yes1 == 1 {
			return 1
		}
		// } else {
		// 	return 0
		// }
	} else {
		*mid = ((*max - 1 - *min) / 2) + *min
		yes2 := searchIndex(index, max, min, mid, inputValue, daftar)
		if yes2 == 1 {
			return 1
		}
	}
	return 0
}

func searchIndex(index *int, max *int, min *int, mid *int, inputValue *int, daftar *[10]int) int {
	if *inputValue == daftar[*mid] {
		*index = *mid
		return 1
	} else if *inputValue > daftar[*mid] {
		*min = *mid + 1
		return 0
	} else if *inputValue < daftar[*mid] {
		*max = *mid - 1
		return 0
	}
	return 0
}
