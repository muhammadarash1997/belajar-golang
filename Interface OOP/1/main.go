// LINK PENJELASAN MANFAAT MENGGUNAKAN INTERFACE
// https://qvault.io/golang/golang-interfaces/#small

package main

import (
	"fmt"
)

type Sport interface {
	sportName() string // artinya fungsi sportName() mempunyai tipe data kembalian string
}

type Human struct {
	name  string
	sport string
}

func (h Human) sportName() string {
	return h.name + " plays " + h.sport + "."
}

func main() {
	human1 := Human{"Rahul", "chess"}
	fmt.Println(human1.sportName())
	human2 := Human{"Riya", "carrom"}
	fmt.Println(human2.sportName())
}
