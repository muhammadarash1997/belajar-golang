// MEMBACA ISI FILE

package main

import (
	"fmt"
	"io"
	"os"
)

var path = "D:/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func readFile() {
	// File yang ingin dibaca harus dibuka terlebih dahulu menggunakan fungsi os.OpenFile() dengan level akses minimal adalah read
	// buka file
	var ourFile, err = os.OpenFile(path, os.O_RDONLY, 0644) // O_RDONLY adalah level akses dimana hanya bisa membaca, 0644 adalah kode permission, berbeda kode dengan kode yg ada di Pembuatan File pun bisa, I don't really know this why
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024) // 10 adalah alokasi elemen yang bisa di tampung, bebas, 10 berarti hanya 10 karakter
	for {
		// fungsi
		n, err := ourFile.Read(text) // err akan menampung antara <nil> (jika tidak ada error) dan error (jika ada error) dan EOF (jika sudah dibaris bawah) n akan menampung angka
		fmt.Println("ini adalah n", n)
		fmt.Println(string(text))
		if err != io.EOF { // EOF adalah sebuah error, biasanya digunakan untuk parameter perbandingan, Error io.EOF sendiri menandakan bahwa file yang sedang dibaca adalah baris terakhir isi atau end of file.
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	// fmt.Println(string(text))
}

func main() {
	readFile()
}

// Error yang muncul ketika eksekusi file.Read() akan di-filter, ketika error tersebut adalah selain io.EOF maka proses baca file akan berlanjut. Error io.EOF sendiri menandakan bahwa file yang sedang dibaca adalah baris terakhir isi atau end of file.
