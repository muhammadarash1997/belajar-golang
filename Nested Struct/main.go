// Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup kompleks
// dengan proses decode hanya sekali

package main

import "fmt"

// NESTED STRUCT SEPERTI DIBAWAH KETIKA KITA INGIN INSTANSIASI MAKA TIDAK BISA LANGSUNG FIELDNYA DIISI
// TETAPI HARUS SATU-PERSATU. BIASANYA DIGUNAKAN UNTUK UNMARSHALL JSON
type person struct {
	name string
	address struct {
		city string
		districts int
	}
}

func main() {
	p1 := person{}

	// CARA INSTANSIASI PERTAMA ADALAH INVALID
	// p1 := person{
	// 	name: "alex",
	// 	address: address{
	// 		city: "Paris",
	// 		districts: "abc",
	// 	},
	// }

	// CARA INSTANSIASI KEDUA ADALAH VALID
	p1.name = "alex"
	p1.address.city = "Paris"
	p1.address.districts = 1

	fmt.Println(p1)
}

