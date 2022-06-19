package main

import "fmt"

// Function
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "GeeksforGeeks" // Memasukkan nilai
	}

	// ini harus pakai close, karena di dalam fungsi main kita menggunakan for
	// dimana di dalamnya ada logic pengecekan yg mana akan berhenti for tsb
	// jika channel close, jika tidak ada logic tsb tidak mengapa tidak
	// menggunakan close
	close(mychnl)
}

// Main function
func main() {

	// Creating a channel
	c := make(chan string)

	// calling Goroutine
	go myfun(c)
	// pada baris ini fungsi myfun masih belum jalan (menunggu data chan masuk)
	for {
		res, ok := <-c // pada baris ini for pada myfun mulai berjalan karena sudah ada yg menerima nilai
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
