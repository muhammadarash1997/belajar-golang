package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println(i)
	i += 5
	return
}

// In this example, the expression “i” is evaluated when the Println call is deferred. The deferred call will print “0” after the function returns.