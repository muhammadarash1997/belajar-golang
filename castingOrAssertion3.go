// Casting Variabel Interface Kosong Ke Objek Pointer

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{} = &person{name: "wick", age: 27}

	fmt.Println(secret.(*person).age)
	fmt.Println(secret.(person).age) // atau bisa juga penulisan castingnya seperti ini
}
