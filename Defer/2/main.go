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
	defer logging() // setelah fungsi runAppilcation() dijalankan maka akan dilanjutkan dengan logging(), harus diletakkan di atas agar tetap dijalankan jika terjadi error
	fmt.Println("Run application")
	result := 10 / value // ketika value yang masuk adalah 0, maka akan terjadi error, tetapi fungsi logging akan tetap dijalankan setelah terjadi error, tetapi akan ditampilkan lebih dulu di console dari pada informasi errornya
	fmt.Println("Result", result)
	// defer logging() <-- letak fungsi yang didefer ada di bawah ini sama saja logging() biasa, artinya kalau di bawah ngapain pakai defer
}
