// Golang program to illustrate
// reflect.Elem() Function

package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Id      int
	Title   string
	Price   float32
	Authors []string
}

// Main function
func main() {
	book := Book{}

	//use of Elem() method
	e := reflect.ValueOf(book)

	fmt.Println(e.NumField())
	fmt.Println(e.Type().Field(0).Name)
	fmt.Println(e.Type().Field(0).Type)
	fmt.Println(e.Field(0).Interface())
	fmt.Println(e.Field(0))

	f := reflect.ValueOf(&book).Elem()

	fmt.Println(f.NumField())
	fmt.Println(f.Type().Field(0).Name)
	fmt.Println(f.Type().Field(0).Type)
	fmt.Println(f.Field(0).Interface())
	fmt.Println(f.Field(0))
}

// reflect.ValueOf().Elem() <-- Digunakan jika yg di reflect adalah alamatnya (metode pointer), alias bukan objectnya, jika yg di reflect adalah alamatnya tetapi tidak menggunakan Elem() maka fungsi2 seperti di bawah ini tidak bisa digunakan
// reflect.ValueOf().Elem().NumField()
// reflect.ValueOf().Elem().Type()
// reflect.ValueOf().Elem().Type().Field()
// reflect.ValueOf().Elem().Type().Field().Name
// reflect.ValueOf().Elem().Type().Field().Type

// reflect.ValueOf() <-- Digunakan jika yg di reflect adalah objectnya (metode biasa), alias alamatnya, jika yg di reflect adalah objectnya tetapi tidak menggunakan Elem() maka fungsi2 seperti di bawah ini tidak bisa digunakan
// reflect.ValueOf().NumField()
// reflect.ValueOf().Type()
// reflect.ValueOf().Type().Field()
// reflect.ValueOf().Type().Field().Name
// reflect.ValueOf().Type().Field().Type
