// return an anonymous function from another function
package main

import "fmt"

// Returning anonymous function
func cetak() func(i, j string) string {
	ourFunc := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return ourFunc
}

func main() {
	value := cetak() // variable value akan menampung sebuah fungsi atau menjadi fungsi
	fmt.Println(value("Welcome ", "to "))
}
