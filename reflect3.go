package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama   string
	Umur   string
	Alamat string
}

func main() {
	a := Sample{"Eko", "32", "Jakarta"}
	z := reflect.ValueOf(a) // {Eko 32 Jakarta}
	fmt.Println(z)          // {Eko 32 Jakarta}

	b := reflect.TypeOf(a) // main.Sample	<-- ini adalah tipe datanya dari object a
	fmt.Println(b)         // main.Sample	<-- ini adalah tipe datanya dari object a

	// strukturFieldPertama := b.Field(0)
	// strukturFieldKedua := b.Field(1)
	// strukturFieldKetiga := b.Field(2)

	// fmt.Println(strukturFieldPertama) // {Name  string  0 [0] false}	<-- Hasil di console
	// fmt.Println(strukturFieldKedua)   // {Age  string  16 [1] false}	<-- Hasil di console
	// fmt.Println(strukturFieldKetiga)  // {Address  string  32 [2] false}	<-- Hasil di console

	// fmt.Println(strukturFieldPertama.Name) // Nama	<-- Hasil di console
	// fmt.Println(strukturFieldKedua.Name)   // Umur	<-- Hasil di console
	// fmt.Println(strukturFieldKetiga.Name)  // Alamat	<-- Hasil di console

	// fmt.Println(strukturFieldPertama.Type.Name()) // string	<-- Hasil di console
	// fmt.Println(strukturFieldKedua.Type.Name())   // string	<-- Hasil di console
	// fmt.Println(strukturFieldPertama.Type.Name()) // string	<-- Hasil di console

	// fmt.Println(strukturFieldPertama.Tag) // kosong, karena kita tidak menulis tag di field tsb	<-- Hasil di console
	// fmt.Println(strukturFieldKedua.Tag)   // kosong, karena kita tidak menulis tag di field tsb	<-- Hasil di console
	// fmt.Println(strukturFieldPertama.Tag) // kosong, karena kita tidak menulis tag di field tsb	<-- Hasil di console
}

// reflect.TypeOf().Field().Tag.Get()
