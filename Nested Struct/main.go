// Teknik ini biasa digunakan ketika decoding data json (json.Unmarshal) yang struktur datanya cukup kompleks
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
	// CARA INSTANSIASI PERTAMA ADALAH INVALID
	p1 := person{
        name: "alex",
        address: struct {
            city     string
            districts int
        }{
            city:     "A",
            districts: 1,
        },
    }
	fmt.Println(p1)

	// CARA INSTANSIASI KEDUA ADALAH VALID
	p2 := person{}
	p2.name = "mike"
	p2.address.city = "B"
	p2.address.districts = 1

	fmt.Println(p2)
}

