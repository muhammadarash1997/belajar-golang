package main

import "fmt"

func main() {

	data := []interface{}{"Hello", 1, "World", 3}

	for key, value := range data {
		switch value.(type) {
		case string:
			fmt.Println(key, value, "(string)")
		case int:
			fmt.Println(key, value, "(int)")
		case []interface{}:
			fmt.Println(key, value, "(interface)")
		}
	}
}
