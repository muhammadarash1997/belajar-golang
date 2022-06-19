package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) Overwrite() {
	// p1 dan p2 akan di overwrite
	p = person{
		name: "alex",
		age:  25,
	}
	fmt.Println(p.name, p.age)
}

func main() {
	var p1 = person{"wick", 21}
	p1.Overwrite()
	var p2 = person{"john", 23}
	p2.Overwrite()
}
