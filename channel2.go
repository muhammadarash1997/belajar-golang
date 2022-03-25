package main

import (
	"fmt"
)

func printChannelData(c chan int) {
	fmt.Println("Data in channel is: ", <-c)
}

func main() {
	fmt.Println("Main started...")
	//create channel of int
	c := make(chan int)
	// call to goroutine
	go printChannelData(c)
	// put the data in channel
	c <- 10
	fmt.Println("Main ended...")
}
