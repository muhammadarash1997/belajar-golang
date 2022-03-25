// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

package main

import (
	"fmt"
)

func printChannelData(nilai chan int) { // goroutine sebenarnya berjalan tetapi dia akan menunggu karena dia belum menerima data channel
	fmt.Println(<-nilai) // program akan terblock/berhenti di sini, akan mulai berjalan jika channel nilai sudah di isi nilai
	for i := 0; i < 5; i++ {
		fmt.Println("test")
	}
}

func main() {
	fmt.Println("Main started...")
	nilai := make(chan int)
	go printChannelData(nilai)

	// saat proses pengiriman nilai ini akan berhenti data menunggu
	// sampai ada yg menerima data, di program ini proses penerimaan
	// data ada di fmt.Println(<-nilai) yg berada di func
	// printChannelData, artinya karena sudah ada yg menerima datanya
	// maka func main akan lanjut menjalankan baris berikutnya
	nilai <- 10
	fmt.Println("Main ended...") // fungsi ini tidak akan dijalankan jika sudah data channel sudah diterima
}

// Intinya Channel digunakan untuk passing nilai dari goroutine ke goroutine, tetap lewat parameter sebuah fungsi tetapi parameter tsb harus bertipe chan
// Jadi jika kita mengirim tanpa ada yg menerima maka akan hang terus, sedangkan jika kita menerima tanpa ada yg mengirim maka akan deadlock
