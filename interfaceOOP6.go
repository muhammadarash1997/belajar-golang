package main

import "fmt"

func main() {
	var eko Person
	eko.Name = "Eko"
	sayHello(eko) // mengapa sayHello bisa diisi objek bertipe Person padahal sayHello meminta yang bertipe HasName, karena kelas Person merupakan turunan karena dia sudah mengimplement atau mewarisi interface HasName

	agus := Person{Name: "Agus"}
	sayHello(agus)
}

type HasName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println(hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string { // karena kelas Person mengimplementasikan fungsi getName() string maka otomatis dia menjadi turunan HasName
	return person.Name
}

// Interface. Jika suatu struct ingin mengimplementasi interface tertentu maka
// struct tersebut wajib memiliki semua fungsi dari interface tsb, tidak bisa
// hanya salah satu fungsi saja.
// Aturan ini berlaku di semua bahasa. Jika suatu class mengimplementasikan
// suatu interface maka class tsb wajib memiliki semua function yang ada di
// interface tsb.