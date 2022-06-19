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
	panic("error occured")
}

// Defer akan tetap dijalankan jika menggunakan panic dan setelah defer dijalankan maka program akan langsung terhenti (exit)
