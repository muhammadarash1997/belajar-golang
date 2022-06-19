package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting application")

	output := flag.Bool("out", false, "Should be there output?")
	input := flag.String("in", "file.csv", "The path to the input file")
	flag.Parse()
	
	fmt.Println(*output)
	fmt.Println(*input)
}