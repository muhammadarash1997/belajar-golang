// Go program to illustrate send
// and receive operation
package main

import (
	"fmt"
	"time"
)

// func myfunc(ch chan int) {

// 	fmt.Println(234 + <-ch)
// }
// func main() { // func main merupakan goroutine
// 	fmt.Println("start Main method")
// 	// Creating a channel
// 	ch := make(chan int)

// 	// golang akan menjalankan func main dan myfunc di thread terpisah karena 22nya merupakan goroutine
// 	go myfunc(ch)
// 	fmt.Println("End Main method")
// 	ch <- 23
// 	time.Sleep(time.Second) // jika tidak diberi jeda maka goroutine func main keburu selesai jadi goroutine myFunc belum sempat menerima data
// }

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() { // func main merupakan goroutine
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	// golang akan menjalankan func main dan myfunc di thread terpisah karena 22nya merupakan goroutine
	go myfunc(ch)
	fmt.Println("End Main method")
	ch <- 23
	time.Sleep(time.Second) // time.Sleep ini mirip delay, jadi program yg ada di bawahnya tidak akan dijalankan jika sampai delaynya selesai, jika tidak diberi jeda maka goroutine func main keburu selesai jadi goroutine myFunc belum sempat menerima data
}
