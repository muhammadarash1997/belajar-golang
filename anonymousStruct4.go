package main

import "fmt"

func main() {
	element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}

	fmt.Println(element)
}
