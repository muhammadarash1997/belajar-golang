package main

import "fmt"

type person struct {
	name string
	age int
}

type child struct {
	name string
	age int
}

type pet struct {
  name string
}

func main() {
	bob := person{name: "bob", age: 15,}
	babyBob := child(bob)
	// "babyBob := pet(bob)" would result in a compilation error

	fmt.Println(bob, babyBob)
}

// An interface type does not have any underlying data structure, but rather exposes a few
// methods of a pre-existing concrete type (which does have an underlying data structure).

// A type assertion brings out the concrete type underlying the interface, while type conversions
// change the way you can use a variable between two concrete types that have the same data structure.