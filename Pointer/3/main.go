// POINTER DI METHOD

package main

import "fmt"

func main() {
	eko := Man{"eko"}
	eko.Married()
	fmt.Println(eko.Name)
}

type Man struct {
	Name string
}

func (man *Man) Married() { // jika tanpa * maka saat print eko.Name di func main, namanya akan tidak berubah
	man.Name = "Mr." + man.Name // menambah Mr di depan namanya
}
