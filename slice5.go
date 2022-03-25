package main

import (
	"fmt"
)

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...) // append(slicePenampung, dataYangInginDitambah)
}

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
	all = RemoveIndex(all, 5)
	fmt.Println(all) //[0 1 2 3 4 6 7 8 9]
}

// a := b[2:4]	<-- artinya a akan menyimpan nilai dari index ke 2 sampai elemen ke 4 nya b
// a := b[:4]	<-- artinya a akan menyimpan nilai dari index ke 0 sampai elemen ke 4 nya b
// a := b[0:4]	<-- artinya a akan menyimpan nilai dari index ke 0 sampai elemen ke 4 nya b
// a := b[:]	<-- artinya a akan menyimpan nilai seluruhnya dari b
// a := b[2:]	<-- artinya a akan menyimpan nilai dari index ke 2 sampai seluruhnya ke belakang b
