package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	p1 := Person{"Arash"}

	f, err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse p1 to byte type so that we can pass p1 as parameter of Write() method
	b, _ := json.Marshal(p1)

	// f is os.File type
	//  So this code below is writing b to myfile.txt
	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}

	// os.Stdout is also has Write() method because it is actually os.File type
	// So this code below is displaying b to console
	_, err = os.Stdout.Write(b)
	if err != nil {
		panic(err)
	}
}
