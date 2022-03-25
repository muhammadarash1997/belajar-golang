package main

import "fmt"

func main() {
	
	var keys = map[int]string{
		1: "aa",
		2: "ab",
		3: "ac",
		4: "ba",
		5: "bb",
		6: "bc",
		7: "ca",
		8: "cb",
		9: "cc",
	}

	for i := 1; i <= len(keys); i++ {
		fmt.Println(keys[i])
	}
}