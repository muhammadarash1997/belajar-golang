package main

import (
	"fmt"
)

func main() {
	newMap := MapBaru("agus")
	fmt.Println(newMap)
}

func MapBaru(name string) map[string]string { // ini adalah sebuah fungsi dengan tipe data kembaliannya adalah map[string]string
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// nil bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
