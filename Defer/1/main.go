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
	// setelah fungsi runAppilcation() dijalankan maka akan dilanjutkan
	// dengan logging(), yang di defer harus diletakkan di atas agar
	// tetap dijalankan jika terjadi error
	// bisa juga diletakkan di bawah, akan tetapi jika fmt.Println("Run application")
	// dijalankan dan terjadi error maka logging tidak akan dijalankan karena belum terbaca
	defer logging()
	fmt.Println("Run application")
	// defer logging()	<-- jika di letakkan di bawah sini maka logging tidak akan dijalankan jika syntax di atasnya terjadi error karena belum terbaca dan juga sama aja seperti logging() biasa artinya kalau dibawah ngapain pakai defer
}
