package main

import "fmt"

type Larik []int

// fungsi append tidak akan berfungsi di dalam method untuk suatu this, akan berfungsi jika this menggunakan pointer, maka kita harus menambahkan *
func (this *Larik) Insert() {
	*this = append(*this, 5)
}

func main() {
	myLarik := Larik{1, 2, 3, 4}
	fmt.Println(myLarik)

	myLarik.Insert()
	fmt.Println(myLarik)
}
