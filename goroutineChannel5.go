package main

import (
	"fmt"
)

func printChannelData(nilai chan int) {
	// close(nilai)	<-- Jika kita pasang close di sini maka (pada syntax di baris 31) var res
	// akan menampung alamat karena channel nilai belum diisi, dan var ok akan menampung
	// false karena sudah di close terlebih dahulu
	fmt.Println("Hello")
	nilai <- 12

	// tidak pakai close juga tidak mengapa, berfungsi agar tidak ada
	// lagi pemasukkan data, agar tidak full memori
	close(nilai)
}

func main() {
	nilai := make(chan int)
	go printChannelData(nilai) // variabel biasa tidak bisa dijadikan argument pada goroutine, maka dari itu ini harus menggunakan channel

	fmt.Println("Main started...")
	fmt.Println("Main ended...")
	fmt.Println(nilai)

	// syntax di bawah ini adalah untuk mengecek apakah channel nilai open atau close, var ok antara true (open) atau false(close), var res akan menampung isi nilainya

	// pada saat inilah fungsi printChannelData mulai menjalankan baris berikutnya
	// yaitu close(nilai), sebelumnya dia menunggu karena tidak ada data chan yg masuk
	res, ok := <-nilai

	fmt.Println(res)
	fmt.Println(ok)
	// nilai <- 32 // <-- ini tidak bisa dilakukan karena channel nilai sudah di close
}

// Intinya Channel digunakan untuk passing nilai dari goroutine ke goroutine, tetap lewat parameter sebuah fungsi tetapi parameter tsb harus bertipe chan
