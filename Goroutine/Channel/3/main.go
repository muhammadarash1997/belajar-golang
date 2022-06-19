// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

package main

import (
	"fmt"
	"time"
)

func printData(c chan *int) {
	time.Sleep(time.Second * 3)
	data := <-c
	time.Sleep(time.Second * 3)
	fmt.Println("Data in channel is: ", *data)
}

func main() {
	fmt.Println("Main started...")
	var a = 10
	b := &a
	// create channel
	c := make(chan *int)
	go printData(c)
	fmt.Println("Value of b before putting into channel", *b)
	c <- b
	a = 20
	fmt.Println("Updated value of a:", a)
	fmt.Println("Updated value of b:", *b)
	time.Sleep(time.Second * 5)
	fmt.Println("Main ended...")
}
