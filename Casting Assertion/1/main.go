// Saat sebuah value dimasukkan ke dalam interface{}
// maka tipe datanya berubah menjadi interface{} juga.
// Lalu ? Bagaimana untuk mengembalikan ke tipe data awal?
// contoh
// var apel interface{} = "Hello World""
// var jeruk string = apel <-- maka ini dilarang dan menyebabkan error, dengan itu assertion akan digunakan

package main

import "fmt"

func main() {
	// var apel interface{} = "Hello World"
	// var jeruk string = apel //<-- maka ini dilarang dan menyebabkan error, dengan itu assertion akan digunakan

	// maka harus seperti di bawah ini
	var apel interface{} = "Hello World"
	var jeruk string = apel.(string) //<-- assertion digunakan
	fmt.Println(apel)
	fmt.Println(jeruk)
}
