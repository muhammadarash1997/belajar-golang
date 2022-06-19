package main

import . "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		s := 4
		k := i
		for j := 1; j <= i; j++ {
			Print(k)
			k += s
			s--
		}
		Println("")
	}
}
