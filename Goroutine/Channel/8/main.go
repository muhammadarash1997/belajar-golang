package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go Only(channel)
	time.Sleep(5 * time.Second)
}

// Jika struktur parameternya seperti dibawah ini maka di dalam fungsi ini bisa mengirim atau menerima data
func Only(channel chan string) {
	channel <- "Hello World" // Mengirim data
	data := <-channel        // Menerima data
	fmt.Println(data)
}

// Jika struktur parameternya seperti dibawah ini maka di dalam fungsi ini bisa mengirim data
// func Only(channel chan<- string) {
// 		channel <- "Hello World" // Mengirim data
// 		data := <-channel        // Menerima data <- maka ini tidak bisa atau error
// 		fmt.Println(data)
// }

// Jika struktur parameternya seperti dibawah ini maka di dalam fungsi ini bisa menerima data
// func Only(channel <-chan string) {
// 		channel <- "Hello World" // Mengirim data <- maka ini tidak bisa atau error
// 		data := <-channel        // Menerima data
// 		fmt.Println(data)
// }
