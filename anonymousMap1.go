package main

import "fmt"

func main() {
	
	var keys map[int]string
	keys = make(map[int]string)
	
	keys[1] = "aa"
	keys[2] = "ab"
	keys[3] = "ac"
	keys[4] = "ba"
	keys[5] = "bb"
	keys[6] = "bc"
	keys[7] = "ca"
	keys[8] = "cb"
	keys[9] = "cc"

	for i := 1; i <= len(keys); i++ {
		fmt.Println(keys[i])
	}
}