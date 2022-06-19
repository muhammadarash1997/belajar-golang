package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text1 = "Banana Burger Soup"
	var regex1, err1 = regexp.Compile(`[a-z]+`)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	var res1 = regex1.FindAllString(text1, -3)
	fmt.Println(res1) // [anana urger oup] <-- karena yang dicari hanya huruf kecil

	var text2 = "Banana Burger Soup"
	var regex2, err2 = regexp.Compile(`[A-Z]+`)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	var res2 = regex2.FindAllString(text2, -3)
	fmt.Println(res2) // [B B S] <-- karena yang dicari hanya huruf besar

	var text3 = "Banana Burger Soup"
	var regex3, err3 = regexp.Compile(`[A-B]+`)
	if err2 != nil {
		fmt.Println(err3.Error())
	}
	var res3 = regex3.FindAllString(text3, -3)
	fmt.Println(res3) // [B B] <-- karena yang dicari hanya huruf besar dan dari A sampai B saja
}
