package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama   string `ini tag dari nama`
	Umur   string `ini tag dari umur`
	Alamat string `ini tag dari alamat`
}

func main() {
	a := Sample{"Eko", "32", "Jakarta"}

	// CARA INISIALISASI 1
	//b := reflect.TypeOf(a)

	// CARA INISIALISASI 2
	var b reflect.Type = reflect.TypeOf(a)

	strukturFieldPertama := b.Field(0)
	strukturFieldKedua := b.Field(1)
	strukturFieldKetiga := b.Field(2)

	fmt.Println(strukturFieldPertama) // {Nama  string  0 [0] false}	<-- Hasil di console
	fmt.Println(strukturFieldKedua)   // {Umur  string  16 [1] false}	<-- Hasil di console
	fmt.Println(strukturFieldKetiga)  // {Alamat  string  32 [2] false}	<-- Hasil di console

	fmt.Println(strukturFieldPertama.Name) // Nama	<-- Hasil di console
	fmt.Println(strukturFieldKedua.Name)   // Umur	<-- Hasil di console
	fmt.Println(strukturFieldKetiga.Name)  // Alamat	<-- Hasil di console

	fmt.Println(strukturFieldPertama.Type.Name()) // string	<-- Hasil di console
	fmt.Println(strukturFieldKedua.Type.Name())   // string	<-- Hasil di console
	fmt.Println(strukturFieldPertama.Type.Name()) // string	<-- Hasil di console

	fmt.Println(strukturFieldPertama.Tag) // ini tag dari nama	<-- Hasil di console
	fmt.Println(strukturFieldKedua.Tag)   // ini tag dari umur	<-- Hasil di console
	fmt.Println(strukturFieldPertama.Tag) // ini tag dari alamat	<-- Hasil di console
}

// reflect.TypeOf().Field().Tag.Get()
