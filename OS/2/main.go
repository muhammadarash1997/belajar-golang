package main

import (
	"fmt"
	"os"
)

var path = "D:/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) { // pengecekan apakah ada filenya atau belum dengan menggunakan trik dengan memasukkan obj error
		var ourFile, err = os.Create(path)
		if isError(err) {
			return
		}
		defer ourFile.Close() // ini untuk menutup filenya karena file yg baru terbuat statusnya otomatis open, defer ourFile.Close() ini akan dijalankan setelah fmt.Println dibawah ini
	}

	fmt.Println("==> file berhasil dibuat", path) // ini akan dijalankan terlebih dahulu setelah itu defer ourFile.Close() akan dijalankan
}

func main() {
	createFile()
}
