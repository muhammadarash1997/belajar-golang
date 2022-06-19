package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// string.NewReader will create strings.Reader object which is
	// strings.Reader is io.Reader type because it implements Read()
	// method of io.Reader interface
	reader := strings.NewReader("hello my name is arash")
	b := make([]byte, 4)

	for {
		_, err := reader.Read(b)	// if we read repeatedly then the contents of b will be overwritten
		if err != nil {

			// io.EOF will come up if data runs out
			if err == io.EOF {
				fmt.Println("data runs out")
				fmt.Println(string(b))
				return
			}
			panic(err)
		}
		fmt.Println(string(b))
	}
}

// OUTPUT OF THIS PROGRAM WILL BE LIKE THIS BELOW

// hell
// o my
//  nam
// e is
//  ara
// shra
// data runs out
// shra