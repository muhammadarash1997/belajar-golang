// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go printDataReceived(messages)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("send data", i)
			messages <- i
		}
	}()

	fmt.Scanln()
}

func printDataReceived(messages chan int) {
	for {
		i := <-messages
		fmt.Println("receive data", i)
	}
}

// Pada func() yang bertugas mencetak receive data kita tidak menggunakan parameter
// channel pada fungsi tsb, karena variabel channel message berada dalam scope yang sama.
// Apabila func() tsb berada di luar func main() atau berada di scope yang berbeda dengan
// variabel message, maka pada func() harus menerima parameter bertipe channel.