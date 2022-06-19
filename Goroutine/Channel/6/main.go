package main

import "fmt"

func main() {

	// Creating a channel
	// Using make() function
	mychnl := make(chan string)

	// Anonymous goroutine
	go func(mychnl chan string) {
		fmt.Println("hey")

		// tidak akan dijalankan jika baris ke 25 (for res := range mychnl) tidak dijalankan, karena fungsi ini menunggu ada yg menerima datanya
		mychnl <- "1"
		mychnl <- "2"
		mychnl <- "3"
		mychnl <- "4"
		close(mychnl)
	}(mychnl)

	fmt.Println("hallo") // akan menampilkan alamat
	// Using for loop
	// for res := range mychnl {
	// 	fmt.Println(res)
	// }
}
