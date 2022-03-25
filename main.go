package main

import (
	"fmt"
)

func main() {
	// magazine := []string{"two", "times", "three", "is", "not", "four"}
	// note := []string{"two", "times", "two", "is", "four"}
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today", "as", "add", "book", "ya"}

	var count int

	func checkMagazine(magazine []string, note []string) {
		// Write your code here
		var count int
		for i := 0; i < len(note); i++ {
			for j := 0; j < len(magazine); j++ {
				if note[i] == magazine[j] {
					count++
					magazine = append(magazine[:j], magazine[(j+1):]...)
					break
				}
			}
		}
		
		if count == len(note) {
			fmt.Print("Yes")
			return
		}
		if count != len(note) {
			fmt.Print("No")
			return
		}
	}
}