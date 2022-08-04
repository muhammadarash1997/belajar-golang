package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go delayer(c)
	for i := 0; i < 3; i++ {
		go receiver(c)
		go transmitter(c)
	}
	fmt.Println("main", <-c)
}

func delayer(c chan int) {
	time.Sleep(5 * time.Second)
	c <- 100
	fmt.Println("delayer")
}

func receiver(c chan int) {
	fmt.Println("receiver", <-c)
}

func transmitter(c chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println("transmitter")
	c <- 50
}