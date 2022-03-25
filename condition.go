package main

import (
	"fmt"
)

func main() {
	a := 7
	b := 6
	if a > b {
		fmt.Println("Hello")
		return // jika ada return di sini maka program yang di bawah fungsi if tidak akan dijalankan karena ada return sehingga fungsi main berhenti
	}
	fmt.Println("Hay")
}
