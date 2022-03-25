package main

import (
	"fmt"
	"os"
)

var path = "D:/test.txt"

func main() {
	var ourFile, err1 = os.Stat(path)
	fmt.Println(err1)
	fmt.Println(ourFile)
	fmt.Println(os.IsNotExist(err1))

	var fileKita, err2 = os.Create(path)
	defer fileKita.Close()
	fmt.Println(err2)
	fmt.Println(fileKita)
	fmt.Println(err2 == nil)

	filePtr, err3 := os.Open("D:/test.txt")
	fmt.Println(err3)
	fmt.Println(filePtr)
}

// ourFile, err1 = os.Stat(path) <-- untuk mengecek apakah sudah ada file yg dimmaksud atau belum, dan akan mengembalikan 2 nilai dan akan ditampung di variabel err1 dan ourFile

// BELUM ADA
// Ketika dicek dengan os.Stat(path) dan ternyata BELUM ADA maka akan mengembalikan 2 nilai yaitu
// err1 = CreateFile D:/test.txt: The system cannot find the file specified. <-- ada error karena belum ada filenya ketika dicek
// ourFile = <nil>
// os.IsNotExist(err1) = true <-- true jika ada error, atau jika error tidak <nil>

// SUDAH ADA
// Ketika dicek dengan os.Stat(path) dan ternyata SUDAH ADA maka akan mengembalikan 2 nilai yaitu
// err = <nil> <-- ga ada error karena sudah ada filenya ketika dicek
// ourFile = &{test.txt 32 {3286044977 30922069} {3286044977 30922069} {3286044977 30922069} 0 0 0 0 {0 0} D:/test.txt 0 0 0 false}
// os.IsNotExist(err1) = false <-- false jika tidak ada error, atau jika error <nil>

// fileKita, err2 = os.Create(path) <-- untuk membuat file, dan akan memiliki 2 nilai kembalian, jika tidak ditampung pun tidak mengapa jika tidak mau memakai nilai kembalian, jika path dan file nya sama dan dibuat ulang maka akan di replace

// KETIKA PATH JALURNYA TIDAK DITEMUKAN DAN ARTINYA GAGAL MEMBUAT FILE
// open D:/temp/test.txt: The system cannot find the path specified.
// <nil>
// (err2 == nil) = false

// KETIKA PATH JALURNYA DITEMUKAN DAN BERHASIL MEMBUAT FILE
// Nilai kembalian akan tetap sama walaupun file yang dimaksud sudah ada
// err2 = <nil>
// fileKita = &{0xc0000d0780}
// (err2 == nil) = true

// os.Open() <-- untuk membuka file

// BELUM ADA
// err3 = open ./test/creating.txt: The system cannot find the path specified.
// filePtr = <nil>

// SUDAH ADA
// err3 = <nil>
// filePtr = &{0xc0000d0780}
