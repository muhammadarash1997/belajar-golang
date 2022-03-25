// MEMBACA ISI FILE

// Jika kita menginginkan sebuah fungsi memiliki parameter parent class/struct, maka kita
// perlu menambah interface untuk bantuan.

// Ada 2 model inheritance di golang yaitu model Inheritance 1 dan Inheritance 2.
// Saat ini kita membahas Inheritance 2.
// Inheritance 2 adalah yang menggunakan bantuan interface.
// Inheritance 2 digunakan ketika ada fungsi manapun yang menerima objectnya child class tetapi
// parameter fungsi tsb bertipe parent classnya.
// Jika ingin seperti itu maka kita butuh bantuan interface

package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct { // Person mengimplementasi interface HasName
	name string
}

func (person Person) GetName() string {
	return person.name
}

type Animal struct { // Animal adalah turunan Person sehingga juga merupakan turunan dari interface HasName sehingga bisa masuk ke fungsi SayHello
	Person
}

func main() {
	var eko Person
	eko.name = "mrEko"
	SayHello(eko)

	var cat Animal
	cat.name = "mrCat"
	SayHello(cat)
}

// Jika kita menginginkan sebuah fungsi menerima parameter bertipe parent class, maka
// tidak cukup dengan model Inheritance 1, kita harus menggunakan model Inheritance 2
// di mana kita harus menggunakan interface sebagai jembatannya seperti di bawah ini.
func SayHello(hasName HasName) {
	fmt.Println(hasName.GetName())
}

// CODE DI BAWAH AKAN ERROR KARENA MASIH MENERAPKAN INHERITANCE 1
// func SayHello(hasName Person) {
// 	fmt.Println(hasName.GetName())
// }