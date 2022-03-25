package main

import (
	"fmt"
)

func main() {
	name := "Mike"
	greeting := "Hello"

	cetak(&name, &greeting)
	fmt.Println(name, greeting)
}

func cetak(name, greeting *string) { //	tipe data name dan greeting sama
	*name = "Ted" // mengakses isi pointer tanpa * akan menampilkan alamatnya, kecuali variabel tsb adalah object maka tanpa * pun bisa
	*greeting = "Hay"
	fmt.Println(*name, *greeting)
}
