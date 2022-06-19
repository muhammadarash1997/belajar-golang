package main

import "fmt"

type Interface interface{}

type Struct struct{}

func main() {
	var a Interface
	fmt.Println(a) // <nil>

	var b Struct
	fmt.Println(b) // {}

	var c string
	fmt.Println(c) // ""

	var d int
	fmt.Println(d) // 0

	var e float64
	fmt.Println(e) // 0

	var f bool
	fmt.Println(f) // false

	var g []int
	fmt.Println(g) // []

	var h map[int]int
	fmt.Println(h) // map[]

	i := &Struct{}
	fmt.Println(i) // &{}

	i = nil
	fmt.Println(i) // <nil>
}
