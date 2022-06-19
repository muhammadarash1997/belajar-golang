// INHERITANCE

// Ada 2 model inheritance di golang yaitu model Inheritance 1 dan Inheritance 2.
// Saat ini kita membahas Inherintace 1.
// Inheritance 1 adalah yang tidak menggunakan bantuan interface dan lebih sederhana.
// Inheritance 1 digunakan ketika tidak ada fungsi manapun yang menerima objectnya
// child class tetapi parameter fungsi tsb bertipe parent classnya.
// Jika ingin seperti itu maka kita butuh bantuan interface.

package main

import "fmt"

type Person struct {
	name string
}
type Book struct {
	Person
}

func (book Person) cetak() {
	fmt.Println("Hello")
}

func main() {
	var agus = Book{}
	agus.name = "arash"
	agus.cetak()
}
