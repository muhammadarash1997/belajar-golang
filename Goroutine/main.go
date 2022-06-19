// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

package main

import (
	"fmt"
	"runtime"
	"time"
)

// syntax yang tidak ada go di depannya, maka meraka dijalankan di thread utama dan mereka dijalankan secara berurutan
func main() {
	runtime.GOMAXPROCS(6) // ini adalah untuk menentukan berapa core yang akan bekerja untuk program ini
	start := time.Now()   // ini hanya untuk mengukur waktu yang dibutuhkan untuk menjalankan program dari awal sampai akhir program
	go func() {           // dengan menambahkan go di depannya, maka dia akan berjalan di lain thread
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	go func() { // dengan menambahkan go di depannya, maka dia akan berjalan di lain thread, fungsi inipun beda thread dengan fungsi go di atas
		for i := 5; i < 8; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start) // ini hanya untuk mengukur waktu yang dibutuhkan untuk menjalankan program dari awal sampai akhir program

	fmt.Println("Total Time For Execution: " + elapsedTime.String()) // ini hanya untuk mengukur waktu yang dibutuhkan untuk menjalankan program dari awal sampai akhir program

	time.Sleep(time.Second) // pada saat program menjalankan baris ini maka time.Sleep akan dijalankan dan ini adalah fungsi untuk menghentikan program, ini hanya untuk menjeda program untuk berhenti, time.Second = 1 detik, time.Hour = 1 jam
	// Jika ada channel, artinya disana ada pengiriman atau penerimaan data antar goroutine maka akan ditunggu
	// Karena goroutine2 di atas tidak berhubungan atau tidak ada channel di sana maka tidak akan saling tunggu
}
