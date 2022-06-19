package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)
}

// perhatikan pada fungsi sayHelloTo kita tidak memerlukan parameter bertipe channel, karena
// variabel channel message berada pada scope yang sama, apabila fungsi sayHelloTo berada di luar
// maka pada fungsi sayHelloTo perlu menerima parameter bertipe channel agar bisa menggunakan
// variabel message