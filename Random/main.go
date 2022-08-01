package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(80))   // return random between 0 until 80
	fmt.Println(rand.Intn(1e6))  // return random between 0 until 1*10^6 atau 1000000
	fmt.Println(rand.Intn(10e6)) // return random between 0 until 10*10^6 atau 10000000
}
