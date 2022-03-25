// DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD VAR

package main

type person struct {
	name string
	age  int
}

func main() {
	// DEKLARASI ANONYMOUS STRUCT DENGAN VAR
	var student1 struct {
		person
		grade int
	}
	student1.person = person{"wick", 21}
	student1.grade = 2

	// KELEMAHAN DEKLARASI ANONYMOUS STRUCT MENGGUNAKAN KEYWORD VAR ADALAH TIDAK BISA LANGSUNG DIINISIALISASI SEPERTI DI BAWAH INI, DI BAWAH INI AKAN EROR
	var student2 struct {
		person
		grade int
	}{
		person = person{"wick", 21}
		grade = 2
	}
}
