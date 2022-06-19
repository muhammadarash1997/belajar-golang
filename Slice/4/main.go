package main

import (
	"fmt"
)

func main() {
	myArray := []int{10, 11, 12, 13, 14, 15}
	fmt.Println("myArray", myArray)

	slice1 := myArray
	fmt.Println("slice1", slice1)

	slice2 := slice1[2:4]
	fmt.Println("slice2", slice2)
	slice2[0] = 1
	fmt.Println("slice2", slice2)

	fmt.Println("myArray", myArray)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	fmt.Println("length of slice1", len(slice1))
	fmt.Println("length of slice2", len(slice2))
}

// Jika slice dicopy maka maka hasil copyan nya adalah reference,
// jadi ketika hasil copyan nya diubah maka slice asli juga akan terubah.

// Jika array dicopy maka hasil copyan nya bukan reference dan hasil copyan nya akan menjadi slice,
// jadi ketika hasil copyan nya diubah maka array asli tidak akan berubah.
// Jika hasil copyan pertama tadi dicopy lagi maka hasil copyan nya adalah reference,
// karena copyan pertama tadi adalah slice