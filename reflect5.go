// PENGGUNAAN reflect.Tag.Get() <-- reflect sendiri bisa diganti variabel hasil reflectnya misal c.Tag.Get()

package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	ID      string `auto_increment:"true" increment:"1"`
	Name    string `varchar: "255"` // <-- jika di Get maka tidak akan tampil karena antara (:) dengan ("255") ada spasi
	Surname string `"varchar: "255"`
}

// Main function
func main() {
	e := Employee{}
	// c variable represents table columns
	c := reflect.TypeOf(e).Field(0).Tag // akan mengambil tag-nya
	fmt.Println(c)

	// // // use of Get method
	// g := c.Get("auto_increment")
	// h := c.Get("increment")
	// fmt.Println(g)
	// fmt.Println(h)

	fmt.Println(reflect.ValueOf(e).Field(0).Interface())
}

// reflect.TypeOf().Field().Tag.Get()
