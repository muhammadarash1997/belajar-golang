package main

import "fmt"

func main() {

	// ARRAY
	odd := [7]int{1, 3, 5, 7, 9, 11, 13}
	for a, b := range odd { // a akan menangkap index sedangkan b menangkap nilainya
		fmt.Println(a, b)
	}

	// STRING
	var apel string = "GeeksforGeeks"
	for i, item := range apel {
		fmt.Printf("string[%d] = %d \n", i, item)
	}

	// MAPS
	jeruk := map[string]int{"Nidhi": 3, "Nisha": 2, "Rohit": 1}
	for c := range jeruk {
		fmt.Println("Rank of", c, "is: ", jeruk[c])
	}
	for c, d := range jeruk {
		fmt.Println("Rank of", c, "is: ", d)
	}
}
