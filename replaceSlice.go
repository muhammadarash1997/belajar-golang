package main

import (
	"fmt"
)

type Mahasiswa struct {
	ID   int
	Name string
	Age  int
}

var DataMahasiswa []Mahasiswa = []Mahasiswa{
	{ID: 1, Name: "Arash", Age: 24},
	{ID: 2, Name: "Agus", Age: 25},
	{ID: 3, Name: "Dodi", Age: 26},
}

func main() {
	fmt.Println(DataMahasiswa)
	DataMahasiswa[1] = Mahasiswa{ID: 5, Name: "Udin", Age: 11}
	fmt.Println(DataMahasiswa)
}
