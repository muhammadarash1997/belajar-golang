package main

import (
	"fmt"
	"runtime"
)

func printMessage(message chan string) {
	fmt.Println(<-message)
}

// Jika fungsi printMessage seperti di bawah ini maka tidak akan bisa karena variabel message berada di scope lain
// func printMessage() {
// 	fmt.Println(<- message)
// }

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	for _, each := range []string{"wick", "hunt", "bourne"} {
	go func(who string) {
		var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}