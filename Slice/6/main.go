package main

import (
	"fmt"
)

func RemoveIndex1D(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}

func RemoveIndex2D(s [][]int, index int) [][]int {
	s = append(s[:index], s[index+1:]...) // append(slicePenampung, dataYangInginDitambah)

	for i := 0; i < len(s[0]); i++ {
		s[i] = RemoveIndex1D(s[i], index)
	}
	return s
}

func main() {
	all := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(all)
	all = RemoveIndex2D(all, 2)
	fmt.Println(all)
}

// a := b[2:4]	<-- artinya a akan menyimpan nilai dari index ke 2 sampai elemen ke 4 nya b
// a := b[:4]	<-- artinya a akan menyimpan nilai dari index ke 0 sampai elemen ke 4 nya b
// a := b[0:4]	<-- artinya a akan menyimpan nilai dari index ke 0 sampai elemen ke 4 nya b
// a := b[:]	<-- artinya a akan menyimpan nilai seluruhnya dari b
// a := b[2:]	<-- artinya a akan menyimpan nilai dari index ke 2 sampai seluruhnya ke belakang b
