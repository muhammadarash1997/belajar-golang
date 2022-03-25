package main

import (
	"fmt"
)

func main() {
	// CARA 1 PENULISAN POINTER OBJEK --> &hero (*g) *hero <-- INI BISA
	g := &hero{
		name:     "Superman",
		greeting: "Hello",
	}
	(*g).greet()
	fmt.Println(g.name) // jika method pada class hero bukan pointer maka hasil g.name akan berbeda dengan yang di dalam method

	// CARA 2 PENULISAN POINTER OBJEK	--> &hero g *hero <-- INI JUGA BISA
	// g := &hero{
	// 	name:     "Superman",
	// 	greeting: "Hello",
	// }
	// g.greet()
	// fmt.Println(g.name)

	// CARA 3 PENULISAN POINTER OBJEK --> hero g *hero <-- ATAU INI JUGA BISA, INI LEBIH SIMPLE KARENA SUDAH TAU DAN TIDAK PERLU KITA MEN-DEREFERENCE-KAN SECARA MANUAL
	// g := hero{
	// 	name:     "Superman",
	// 	greeting: "Hello",
	// }
	// g.greet()
	// fmt.Println(g.name)

}

type hero struct {
	name     string
	greeting string
}

func (objek *hero) greet() {
	fmt.Println(objek.greeting, objek.name)
	objek.name = "Bagus"
}
