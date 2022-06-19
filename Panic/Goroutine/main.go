package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(6)

	defer fmt.Println("main goroutine ended")
	go secondGoroutines()
	go firstGoroutines()

	time.Sleep(time.Second)
}

func firstGoroutines() {
	defer fmt.Println("first goroutines ended")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	panic("first goroutines error")
}

func secondGoroutines() {
	defer fmt.Println("second goroutines ended")
	for i := 5; i < 20; i++ {
		fmt.Println(i)
	}
}

// Kesimpulannya adalah panic akan menghentikan goroutine lain dan
// akan menjalankan fungsi defer tetapi tidak dengan fungsi defer yang
// ada di goroutine lain