package main

import "fmt"

func main() {
	fmt.Println(c())
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

// In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2