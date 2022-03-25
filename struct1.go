package main

import (
	"fmt"
)

type Book struct {
	name   string
	author string
	pages  int
}

func (book Book) cetak() {
	fmt.Println(book.name, book.author)
	fmt.Println(book.pages)
}

func main() {
	book1 := Book{"Monster Blood", "Roger", 131}
	book1.cetak()
	book1.name = "Vampire Breath"
	book1.pages = 162
	book1.cetak()
}
