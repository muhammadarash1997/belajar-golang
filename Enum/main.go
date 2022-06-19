package main

import (
	"fmt"
	"reflect"
)

type Role int

const (
	Manager    Role = iota // Head = 0
	Supervisor             // Shoulder = 1
	Staff                  // Knee = 2
)

func (r Role) String() string {
	return []string{"Manager", "Supervisor", "Staff"}[r]
}

type Orang struct {
	Name string
	Role Role
}

func main() {
	person1 := Orang{
		Name: "Abdul",
		Role: Manager,
	}
	person2 := Orang{
		Name: "Muchlis",
		Role: Supervisor,
	}

	fmt.Println(person1)
	fmt.Println(person1.Name)
	fmt.Println(person1.Role)
	fmt.Println(reflect.TypeOf(person1.Name))
	fmt.Println(reflect.TypeOf(person1.Role.String()))

	fmt.Println("")

	fmt.Println(person2)
	fmt.Println(person2.Name)
	fmt.Println(person2.Role)
	fmt.Println(reflect.TypeOf(person2))
	fmt.Println(reflect.TypeOf(person2.Name))
	fmt.Println(reflect.TypeOf(person2.Role.String()))
}
