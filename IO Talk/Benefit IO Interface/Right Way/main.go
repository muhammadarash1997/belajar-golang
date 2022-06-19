package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
}

func (p Person) Write(w io.Writer) {
	b, _ := json.Marshal(p)
	
	_, err := w.Write(b)
	if err != nil {
		panic(err)
	}
}

func main() {
	p1 := Person{"Arash"}

	f, err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(f.Name())

	// this is the advantage of io interface we can pass anything
	// has io.Writer type
	p1.Write(f)
	p1.Write(os.Stdout)
}
