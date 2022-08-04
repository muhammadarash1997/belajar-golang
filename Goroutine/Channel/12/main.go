package main

import "fmt"

func main() {
	waitc := make(chan struct{}) // Jadi channel disini bukan untuk mengirim data melainkan untuk dimanfaatkan agar dapat menunggu fungsi goroutine selesai

	go func() {
		var i = 0
		for {
			if i == 10 {
				close(waitc)
				return
			}
			fmt.Println(i)
			i++
		}
	}()
	<-waitc
}

// <-waitc for waiting goroutine done so func main will not end until waitc channel is closed. This is for waiting by utilizing deadlock concept.
