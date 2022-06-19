package main

import "fmt"

type tankOne interface {
	Tarea() float64
}

type tankTwo interface {
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func (m myvalue) Multiply() float64 {
	return m.height * m.radius
}

func main() {
	// var t tankOne
	// t = myvalue{10, 14}
	// fmt.Println("Area of tank :", t.Tarea())
	// fmt.Println("Volume of tank:", t.(tankTwo).Volume())     // Ketika ingin menggunakan fungsi dari interface tankTwo maka objek t harus dalam keadaan bertipe tankTwo
	// fmt.Println("Multiply of tank:", t.(myvalue).Multiply()) // Ketika ingin menggunakan fungsi dirinya sendiri maka objek t harus dalam keadaan bertipe dirinya sendiri

	var t myvalue
	t = myvalue{10, 14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())
	fmt.Println("Multiply of tank:", t.Multiply())
}

// Interface. Jika suatu struct ingin mengimplementasi interface tertentu maka
// struct tersebut wajib memiliki semua fungsi dari interface tsb, tidak bisa
// hanya salah satu fungsi saja.
// Aturan ini berlaku di semua bahasa. Jika suatu class mengimplementasikan
// suatu interface maka class tsb wajib memiliki semua function yang ada di
// interface tsb.
// Suatu struct tsb dapat mengimplementi lebih dari 1 interface hanya saja ketika objek dari
// struct tsb sedang bertipe interface A dan ingin menggunakan fungsi dari interface B maka
// harus dalam keadaan bertipe interface B
// Suatu struct tsb dapat memiliki fungsinya sendiri yang bukan dari interface saja,
// tetapi jika saat ingin menggunakan fungsinya sendiri object tsb
// harus dalam keadaan bertipe dirinya sendiri atau tidak dalam keadaan
// bertipe interface
