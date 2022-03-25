// Casting Variabel Interface Kosong Ke Objek

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{} = person{name: "wick", age: 27}

	fmt.Println(secret.(person).age) // harus di casting dulu ke type data asalnya walaupun langsung di print kalau isinya berupa objek, artinya tidak bisa langsung secret.age

	var abu interface{} = 13
	fmt.Println(abu) // abu bisa langsung di print karena isinya bukan objek

	var apel = abu.(int) // tetapi kalau dia mau di tampung ke variabel yg di tulis type datanya maka harus di casting dulu
	fmt.Println(apel)

	var salak int = abu.(int) // dan juga kalau dia mau di tampung ke tipe lain maka harus di casting dulu
	fmt.Println(salak)

	var jeruk interface{} = abu // kalau penampungnya juga tipe interface maka tidak perlu di casting
	fmt.Println(jeruk)
}
