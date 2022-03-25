// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("End App")
	message := recover() // akan menangkap eror sehingga program tidak jadi berhenti
	fmt.Println("Pesan erornya adalah", message)
}

func runApp(adaEror bool) {
	defer endApp() // tangkap erornya di dalam fungsi endApp dengan menggunakan recover() di dalamnya
	if adaEror {
		panic("ERROR TERJADI BOI") // jadi ketika panic terjadi, maka fungsi yang didefer akan tetap dijalankan, dijalankan terlebih dahulu sebelum panic jika dilihat di console
	}
	// message := recover() <-- jika recover() diletakkan di sini, maka tidak akan menangkap erornya, tetapi seharunya diletakkan di dalam fungsi yang didefer
}
