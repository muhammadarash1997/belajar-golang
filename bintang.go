package main

import . "fmt"

func main() {
	for i := 0; i < 5; i++ {
		Println("")
		for j := 0; j <= i; j++ {
			Print("w")
		}
	}
}
