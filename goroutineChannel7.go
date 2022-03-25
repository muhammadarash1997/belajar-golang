package main

import (
	"fmt"
	"time"
)

func main() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		fmt.Println("fungsi goroutine dijalankan")
		time.Sleep(5000000000)
		ch <- 42
	}()
	// read from channel
	fmt.Println("ini adalah awal")
	fmt.Println("ini adalah awal")
	fmt.Println("ini adalah awal")
	fmt.Println("ini adalah awal")
	fmt.Println("ini adalah awal")
	fmt.Println(<-ch) // ketika program membaca baris ini, barulah go func {} dijalankan, dan setelah selesai dijalankan, barulah baris ini dijalankan
	fmt.Println("ini adalah akhir")
}
