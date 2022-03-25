var agus int       // <-- artinya agus adalah variabel yg bisa menyimpan nilai bertipe int
var agus []int     // <-- artinya agus adalah variabel yg bisa menyimpan nilai bertipe slice yg mana slice bertipe int
var agus struct
var agus = struct{...}{...}	// <-- artinya jika var tidak memiliki tipe data maka harus di isi nilai, maka agus akan menjadi variabel yg memiliki tipe data struct
var agus = []struct{...}{...}	// <-- artinya jika var tidak memiliki tipe data maka harus di isi nilai, maka agus akan menjadi variabel yg memiliki tipe data slice yg mana slice bertipe struct

type agus int      // <-- artinya agus adalah type alias dari int
type agus struct { // <-- artinya agus adalah type alias dari struct
	nama string // <-- artinya nama adalah field yg bisa menyimpan nilai bertipe string
	usia int
}
type agus func(string) string // <-- artinya agus adalah type alias dari funngsi yg memiliki input parameter bertipe string dan output parameter bertipe string
type agus func(*string) // <-- artinya agus adalah type alias dari funngsi yg memiliki input parameter bertipe pointer string

var agus struct {	// <-- artinya agus adalah variabel yg bisa menyimpan nilai bertipe struct
	nama int	// <-- artinya agus adalah field yg bisa menyimpan nilai bertipe string
	usia int
}

agus := 13
agus := struct { // <-- karena menggunakan tanda := maka harus langsung memberi nilai, maka ini juga anonymous
	name	string
	usia	int
}{
	name: "agus",
	usia: 24
}

agus := [6]int{9, 10, 11, 12, 13, 14} // <-- ini sebenarnya juga anonymous slice