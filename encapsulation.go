package person

func Read() { // ini adalah public karena huruf awal besar, artinya bisa diakses dari luar person package
	// implementation
}

func Write() { // ini adalah private function karena huruf awal kecil, artinya hanya tersedia di person package
	// implementation
}

// public struct
type Book struct {
	Name   string // ini adalah public karena huruf awal besar, artinya bisa diakses dari luar person package
	author string // ini adalah private field karena huruf awal kecil, artinya hanya tersedia di person package
}
