package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello World"

	// os.Create will create file and will return os.File object
	// which is this os.File is io.Reader type because os.File implements
	// Read() method of io.Reader interface and os.File is also io.Writer
	// type because os.File implements Write() method of io.Writer interface
	f, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fmt.Println(f.Name())	// will print myfile.txt

	n, err := f.Write([]byte(text))
	if err != nil {
		fmt.Println(err)
		return
	}

	// n will be 11 because "Hello World" is 11 byte
	fmt.Println(n)
}
