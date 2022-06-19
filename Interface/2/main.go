package main

import (
	"fmt"
)

// dengan interface atau tanpa interface sama saja
/*type Animaler interface {
	Eat()
	Move()
	Speak()
	Error()
}*/

type SuperAnimals struct {
	locomotion string
}

type Animals struct {
	SuperAnimals
	food  string
	sound string
}

func (x Animals) Eat() {
	fmt.Println(x.food)
}

func (x Animals) Move() {
	fmt.Println(x.locomotion)
}

func (x Animals) Speak() {
	fmt.Println(x.sound)
}

func (x Animals) Error() {
	fmt.Println("Invalid query entered!")
}

func main() {
	m := Animals{SuperAnimals{"walk"}, "grass", "moo"}
	m.Move()
}
