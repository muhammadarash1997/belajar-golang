package main

import (
	"flag"
	"fmt"
)

const (
	defaultName = "world"
)

var (
	address = flag.String("addr", "localhost:50051", "the address to connect to")	// "addr" is the flag, "localhost:50051" is default value of address variable if addr flag not filled, "the address to connect to" is explanation what for that flag
	theName = flag.String("name", defaultName, "Name to greet")	// "name" is the flag, defaultName is default value of theName, "Name to greet" is explanation what for that flag
)

func main() {
	flag.Parse() // <-- pakai ini atau tidak hasil tetap sama

	fmt.Println(*address) // localhost:50051
	fmt.Println(*theName) // world
}

// go run flag1.go <-- address will contains default value namely "localhost:50051"
// go run flag1.go --addr localhost:8080 <-- address will contains "localhost:8080"

// go run flag1.go <-- theName will contains default value namely "world"
// go run flag1.go --name guys <-- theName will contains "guys"

// TRY THIS BELOW ON COMMAND LINE
// go run flag1.go --addr localhost:8080 --name guys <-- address will contains "localhost:8080" theName will contains "guys"

// TRY THIS BELOW ON COMMAND LINE
// go run flag1.go --help