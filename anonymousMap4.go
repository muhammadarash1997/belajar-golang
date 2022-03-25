package main

import "fmt"

func main() {
	
	type keys map[int]string	// keys merupakan tipe alias dari map[int]string
	a := keys{1: "aa"}	// a merupakan objek baru bertipe keys
	b := keys{1: "bb"}	// b merupakan objek baru bertipe keys
	c := keys{1: "cc"}	// c merupakan objek baru bertipe keys

	fmt.Println(a)	// map[1:aa]
	fmt.Println(a[1])	// aa

	fmt.Println(b)	// map[1:aa]
	fmt.Println(b[1])	// aa

	fmt.Println(c)	// map[1:aa]
	fmt.Println(c[1])	// aa
}