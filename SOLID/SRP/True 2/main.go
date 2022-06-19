package main

import (
	"fmt"
	"log"
	"regexp"
)

type Email struct {
	email string
}

func NewEmail(email string) Email {
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
		log.Fatal("Invalid email")
	}
	return Email{email}
}

func (this *Email) GetEmail() string {
	return this.email
}


type Person struct {
	name    string
	surname string
	email   Email
}

func NewPerson(name string, surname string, email Email) *Person {
	return &Person{name, surname, email}
}

func (this *Person) SetName(name string) {
	this.name = name
}
func (this *Person) SetSurname(surname string) {
	this.name = surname
}
func (this *Person) SetEmail(email Email) {
	this.email = email
}
func (this *Person) GetName() string {
	return this.name
}
func (this *Person) GetSurname() string {
	return this.surname
}
func (this *Person) GetEmail() string {
	return this.email.GetEmail()
}

func CreatePerson() *Person {
	email := NewEmail("naruto@gmail.com")
	person := NewPerson("Naruto", "Uzumaki", email)
	return person
}

func main() {
	thePerson := CreatePerson()
	
	fmt.Println(thePerson.GetName())
	fmt.Println(thePerson.GetSurname())
	fmt.Println(thePerson.GetEmail())
}

// func main() {
// 	email := NewEmail("naruto@gmail.com")
// 	thePerson := NewPerson("Naruto", "Uzumaki", email)

// 	fmt.Println(thePerson.GetName())
// 	fmt.Println(thePerson.GetSurname())
// 	fmt.Println(thePerson.GetEmail())
// }