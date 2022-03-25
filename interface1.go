package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Hello" // type data interface seperti type data dynamic di dart, artinya bisa tipe apa saja bahkan bisa diisi alamat
	i = 13                      // i bisa dirubah menjadi int atau apapun
	switch i.(type) {
	case int:
		fmt.Println("it is int")
	case float64:
		fmt.Println("it is float64")
	case string:
		fmt.Println("it is string")
	default:
		fmt.Println("it is another type")
	}
}
