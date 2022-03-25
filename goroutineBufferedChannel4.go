// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

// Channel Direction

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 0)

	go printDataReceived(messages)

	go sendData(messages)

	fmt.Scanln()
}

func printDataReceived(messages <-chan int) {
	for {
		// variabel messages digunakan untuk mengirim
		i := <-messages
		fmt.Println("receive data", i)
	}
}

func sendData(messages chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		// variabel messages digunakan untuk menerima
		messages <- i
	}
}

// Secara default jika chan tanpa tanda (<-) maka bisa
// digunakan untuk mengirim dan menerima data.
// (messages chan<- int) artinya variabel messages hanya bisa digunakan untuk menerima data.
// (messages <-chan int) artinya variabel messages hanya bisa digunakan untuk mengirim data.