package main

import "fmt"

func main() {
	myInt := IntCounter(6)
	var inc Adder = &myInt // mengapa tipe data Adder bisa menampung objek bertipe IntCounter, karena class IntCounter telah mewarisi atau mengimplement class Adder
	for i := 0; i < 10; i++ {
		fmt.Println(inc.add())
		fmt.Println(myInt)
	}
}

type Adder interface {
	add() int // jika ada kelas yang mempunyai fungsi persis seperti yang dimiliki interface maka otomatis kelas tsb mengimplementasi atau mewarisi interface ini
}

type IntCounter int // tidak hanya tipe data struct yang bisa mempunyai interface, tetapi tipe data apapun bisa

func (obj *IntCounter) add() int { // dengan menambahkan fungsi add() int maka otomatis golang mengetahui bahwa kelas IntCounter mewarisi interface Adder
	*obj++
	return int(*obj)

}