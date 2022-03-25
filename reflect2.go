package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `requiredmaksimal:"10 karakter boi"` // <-- jika ingin bisa dipanggil dengan Get() nya reflect maka key dari tag harus tidak ada spasi, misalnya (required maksimal) maka ini salah, seharusnya (requiredmaksimal)
	Age  int    `hello dunia`                        // <-- pada Tag ini tidak mengandung key dan value, maka tidak bisa menggunakan Get nya reflect
}

func main() {
	a := Animal{"cat", 10}
	t := reflect.TypeOf(a) // t adalah cerminan dari a hanya saja sekarang memiliki banyak fitur fungsi, tetapi yg dicerminkan adalah type datanya
	fmt.Println(t)         // akan mengeluarkan tipe main.Animal di console, mengapa ada main nya dan bukan Animal saja, karena dalam package main
	// FieldByName() akan mengembalikan 2 values maka dari itu harus ada 2 variabel yg menampung
	field, cek := t.FieldByName("Name")            // cek adalah untuk mengecek apakah atribut Age ada di dalam t, nilai yang ditampung cek antara true or false
	fmt.Println(field)                             // {Age  int hello dunia 16 [1] false} <-- hasil yg ada di console
	fmt.Println(cek)                               // true or false
	fmt.Println(field.Tag)                         // hello dunia <-- hasil yg ada di console
	fmt.Println(field.Tag.Get("requiredmaksimal")) //
}

// reflect.TypeOf().Field().Tag.Get()
