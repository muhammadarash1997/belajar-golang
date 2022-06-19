package main

import (
	"fmt"
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
	return
}

// Defer akan tetap dijalanakan jika menggunakan return
