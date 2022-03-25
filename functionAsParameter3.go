package main

import (
	"fmt"
)

type fungsi func(int)

var apel1 = func(a int) { fmt.Println(a) }

func apel2(a int) {
	fmt.Println(a)
}

func main() {
	makan(apel1)
	makan(apel2)
}

func makan(f fungsi) {
	f(12)
}
