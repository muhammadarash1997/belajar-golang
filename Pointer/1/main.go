package main

import "fmt"

func main() {
	var s1 = student{Name: "john", Grade: 2}
	coba(s1, &s1)
	fmt.Println("Alamat asli adalah", &s1.Name) // alamat asli Name
}

type student struct { // kelas student
	Name  string
	Grade int
}

func (obj student) learn() { // methodnya kelas student
	fmt.Printf("%v is learning\n", obj.Name)
}

func coba(object1 student, object2 *student) { // cara biasa akan memakan memori baru
	fmt.Println(object1.Name, object2.Name) // cara pengaksesan sama antara biasa dan pointer
	object1.learn()                         // cara pengaksesan sama antara biasa dan pointer
	object2.learn()                         // cara pengaksesan sama antara biasa dan pointer

	fmt.Println(&object1, object2) // akan menghasilkan output yang sama

	fmt.Println(object1, *object2) // akan menghasilkan output yang sama

	fmt.Println(&object1.Name, &object2.Name) // alamat object pointer akan sama dengan aslinya
}
