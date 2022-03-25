package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("halo")
	os.Exit(4)
	fmt.Println("selamat datang")
}

// Exit digunakan untuk menghentikan program secara paksa pada saat itu juga.
// Semua statement setelah exit tidak akan di eksekusi, termasuk juga defer.
// Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah
// parameter bertipe numerik yang wajib diisi. Angka yang dimasukkan akan
// muncul sebagai exit status ketika program berhenti.
