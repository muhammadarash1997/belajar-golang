package main

import "fmt"

func main() {
	geoHash := "karakter"

	fmt.Println(geoHash)

	for i := 0; i < 5; i++ {
		geoHash = geoHash[:len(geoHash)-1]
		fmt.Println(geoHash)
	}
}
