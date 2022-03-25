package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// Anonymous Struct
	// ini merupakan pembuatan struct sekaligus penciptaan objek
	var s1 = struct { // s1 merupakan sebuat object dari struct
		person
		grade int
	}{}
	// Inisialisasi
	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name :", s1.person.name)
	fmt.Println("age :", s1.person.age)
	fmt.Println("grade :", s1.grade)

	// Anonymous Struct sekaligus Inisialisasi
	// ini merupakan pembuatan struct sekaligus penciptaan objek sekaligus inisialisasi
	var s2 = struct { // <-- struct{} merupakan tipe data, ini sama seperti []int dari var s2 = []int{...}
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	fmt.Println("name :", s2.person.name)
	fmt.Println("age :", s2.person.age)
	fmt.Println("grade :", s2.grade)

}

// var s2 = struct { // <-- struct{} merupakan tipe data, ini sama seperti []int dari var s2 = []int{...}
// 	person
// 	grade int
// }{
// 	person: person{"wick", 21},
// 	grade:  2,
// }

// []int <-- artinya slice bertipe int
// []struct{...} <-- artinya slice bertipe struct{...}
