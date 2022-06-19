package main

import "fmt"

func main() {
	k := make(map[string]int)
	k["aaa"] = 1
	k["iii"] = 2
	k["uuu"] = 3
	k["eee"] = 4
	var count int
	for _, d := range k {
		fmt.Println(d)
		count++
		if count == 4 {
			k["ooo"] = 5
			for _, e := range k {
				fmt.Println(e)
			}
		}
	}
}
