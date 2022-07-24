package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwordHash1, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)
	passwordHash2, _ := bcrypt.GenerateFromPassword([]byte("arashpassword"), bcrypt.MinCost)

	fmt.Println(passwordHash1)
	fmt.Println(passwordHash2)

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash2), []byte("arashpassword"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
