package main

import (
	"fmt"
	"strings"
)

func main() {
	// string.NewReader will create strings.Reader object which is
	// strings.Reader is io.Reader type because it implements Read()
	// method of io.Reader interface
	reader := strings.NewReader("hello my name is arash")
	b := make([]byte, 4)

	// Read will store data read to b and
	// return number of data bytes read but
	// if data read is larger than b size
	// then it will return max of b can contain
	// or in other words is the size of b
	n, err := reader.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(b))
}

// OUTPUT OF THIS PROGRAM WILL BE LIKE THIS BELOW

// OUTPUT OF THIS PROGRAM WILL BE LIKE THIS BELOW

// 4
// hell