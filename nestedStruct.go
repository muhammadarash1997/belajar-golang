// Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup kompleks
// dengan proses decode hanya sekali

package main

func main() {
}

type student struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}
