package main

import "fmt"

func main() {
	// MEMBUAT OBJEK DENGAN CARA 1
	var kecoa *Monster
	kecoa = &Monster{32}
	fmt.Println((*kecoa).usia) // (*kecoa) <-- mengakses nilai pointer maka harus menggunakan *, jika tidak akan keluar alamatnya
	fmt.Println(kecoa.usia)    // sebenarnya untuk mengakses bisa langsung seperti ini, jadi tidak usah seperti (*kecoa) karena golang sudah mengerti

	// MEMBUAT OBJEK DENGAN CARA 2
	var ular *Monster
	fmt.Println(ular)   // pointer yang belum ada isinya akan berisi <nil>
	ular = new(Monster) // cara membuat objek dengan syntax new tidak bisa langsung mengisi field, usia default akan 0 karena type data int
	(*ular).usia = 42   // (*ular) <-- untuk mengakses nilai pointer maka harus menggunakan *, jika tidak akan keluar alamatnya
	fmt.Println((*ular).usia)
	fmt.Println(ular.usia) // sebenarnya untuk mengakses bisa langsung seperti ini, jadi tidak usah seperti (*ular) karena golang sudah mengerti
}

type Monster struct {
	usia int
}
