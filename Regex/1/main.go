package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`) // Ekspresi [a-z]+ maknanya adalah, semua string yang merupakan alphabet yang hurufnya kecil

	if err != nil {
		fmt.Println(err.Error())
	}

	// Regex.FindAllString(text, -1)	<-- text merupakan ekpresinya, lalu angka -1 adalah jumlah hasil pencarian
	// Jika batas diset -1, maka akan mengembalikan semua data, diset -2 pun mengembalikan semua nilai, pokoknya kalau nilai min maka akan mengembalikan semua nilai

	var res = regex.FindAllString(text, -3) // Jika diset min maka akan mengembalikan semua data
	fmt.Println(res)                        // [banana, burger, soup]

	var res0 = regex.FindAllString(text, -2) // Jika diset min maka akan mengembalikan semua data
	fmt.Println(res0)                        // [banana, burger, soup]

	var res1 = regex.FindAllString(text, -1) // Jika diset min maka akan mengembalikan semua data
	fmt.Println(res1)                        // [banana, burger, soup]

	var res2 = regex.FindAllString(text, 0)
	fmt.Println(res2) // []	<-- menampilkan 0 buah nilai karena memang yg diminta 0

	var res3 = regex.FindAllString(text, 1)
	fmt.Println(res3) // [banana]	<-- menampilkan 1 buah nilai karena memang yg diminta 1

	var res4 = regex.FindAllString(text, 2)
	fmt.Println(res4) // [banana, burger]	<-- menampilkan 2 buah nilai karena memang yg diminta 2

	var res5 = regex.FindAllString(text, 3)
	fmt.Println(res5) // [banana, burger, soup]	<-- menampilkan 3 buah nilai karena memang yg diminta 3

	var res6 = regex.FindAllString(text, 4)
	fmt.Println(res6) // [banana, burger, soup]	<-- mengapa menampilkan 3 buah nilai saja padahal diminta 4, karena data yang tersedia hanya 3
}
