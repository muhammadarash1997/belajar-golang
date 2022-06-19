package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number) // baris ini seperti membuat copian dari variabel number, jadi reflectValue adalah copian number tetapi sekarang memiliki fungsi2 yang ada di dalam reflect
	fmt.Println(reflectValue)                  // akan menghasilkan isi number karena ValueOf() akan mencerminkan Valuenya, sedangkan TypeOf() akan mencerminkan Typenya

	// reflectValue.Type() dan reflectValue.Kind() akan menghasilkan output yang sama ketika diprint tetapi jika untuk comparison kita gunakan .Kind()
	fmt.Println(reflectValue.Type())
	if reflectValue.Kind() == reflect.Int { // reflect.Int merupakan dari package reflect langsung
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}

// Untuk Object bisa menggunakan ValueOf dan TypeOf
// Untuk Variable biasa juga bisa menggunakan ValueOf dan TypeOf
// reflect.ValueOf()
// reflect.ValueOf().Type()
// reflect.ValueOf().Kind()
// reflect.ValueOf().Int()
// reflect.TypeOf() contoh -->	a := reflect.TypeOf(b) <-- misal tipe data b adalah int, maka nilai yg ditampung a adalah int, jika tipenya class maka yg ditampung a adalah tipe class
// reflect.TypeOf().Field() contoh --> a := reflect.TypeOf(b).Field(0)
// reflect.TypeOf().FieldByName() <-- akan mengembalikan 2 buah nilai
// reflect.TypeOf().Field().Tag	<-- Tag biasanya digunakan jika yg di reflect adalah variabel yg tipe datanya class, karena tag hanya ada di field dan field hanya ada di class
// reflect.TypeOf().Field().Tag.Get() <-- Get digunakan untuk mengambil isi Tag nya berdasarkan key dari Tag, jadi maksudnya jika di dalam Tag-nya berupa key dan value, maka Get ini biasanya untuk mengambil value di dalam Tag tsb dengan memasukkan Key ke Get
