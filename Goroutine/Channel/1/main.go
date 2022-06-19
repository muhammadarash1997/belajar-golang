// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

package main

import (
	"fmt"
)

func printChannelData(nilai int) {
	fmt.Println(nilai)

}

func main() {
	nilai := 10
	go printChannelData(nilai) // variabel biasa tidak bisa dijadikan argument pada goroutine, maka dari ini harus menggunakan channel
	fmt.Println("Main started...")
	fmt.Println("Main ended...")
}

// Intinya Channel digunakan untuk passing nilai dari goroutine ke goroutine, tetap lewat parameter sebuah fungsi tetapi parameter tsb harus bertipe chan
