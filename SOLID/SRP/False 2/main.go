package main

import (
	"fmt"
	"log"
	"regexp"
)

type Person struct {
	name    string
	surname string
	email   string
}

func NewPerson(name string, surname string, email string) *Person {
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
		log.Fatal("Invalid email")
	}
	return &Person{name, surname, email}
}

func main() {
	thePerson := NewPerson("Naruto", "Uzumaki", "naruto@gmail.com")
	fmt.Println(thePerson.name)
	fmt.Println(thePerson.surname)
	fmt.Println(thePerson.email)
}