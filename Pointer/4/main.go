package main

import "fmt"

func main() {
	a := 8
	b := &a
	c := &b
	d := &c

	fmt.Println(*b)

	fmt.Println(a) // 8
	fmt.Println(b) // 0xc000012078 <-- alamat dari variabel a
	fmt.Println(c) // 0xc000006028	<-- alamat dari variabel b
	fmt.Println(d) // 0xc000006030	<-- alamat dari variabel c

	// fmt.Println(*a)	// ini tidak bisa, karena a tidak menyimpan alamat
	fmt.Println(*b) // 8 <-- isi dari alamat yg disimpan variabel b yaitu a
	fmt.Println(*c) // 0xc000012078 <-- alamat dari variabel a
	fmt.Println(*d) // 0xc000006028 <-- alamat dari variabel b

	fmt.Println(a)    // 8
	fmt.Println(*b)   // 8
	fmt.Println(**c)  // 8
	fmt.Println(***d) // 8
}
