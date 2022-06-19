package main

import "fmt"

func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}

// Output :
/// halo 1
/// halo 3
/// halo 2

// halo 3 lebih dulu daripada halo 2 karena halo 3 berada di dalam fungsi sendiri
