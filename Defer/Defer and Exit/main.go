package main

import (
	"fmt"
	"os"
)

func main() {
	runApplication(0) // mengirim nilai
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	os.Exit(0)
}

// Defer tidak akan dijalankan jika ada os.Exit
