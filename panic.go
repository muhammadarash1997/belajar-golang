// Panic function adalah function yang bisa kita gunakan untuk menghentikan program, artinya ketika panic dijalankan maka program akan berhenti, tetapi bisa juga tetap dijalankan dengan cara menambahkan fungsi recover()
// Panic function biasanya dipanggil ketika terjadi error pada saat program berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi

package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("Apakah ini dijalankan") // baris ini tidak akan dijalankan karena program telah berhenti
}

func endApp() {
	fmt.Println("End App")
}

func runApp(adaEror bool) {
	defer endApp()
	if adaEror {
		panic("ERROR TERJADI BOI") // jadi ketika panic terjadi, maka fungsi yang didefer akan tetap dijalankan, dijalankan terlebih dahulu sebelum panic jika dilihat di console
	}
}

// Biasanya panic digunakan pada untuk pengaplikasian seperti,
// jika program sedang berjalan dan program sedang mencoba untuk
// connect ke database misalnya, dan ternyata tidak bisa, maka
// kita bisa gunakan/tambahkan panic sehingga program berhenti
