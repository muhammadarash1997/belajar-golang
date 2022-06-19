package main

import (
	"fmt"
	"io"
)

var path = "D:/test.txt"

func main() {
	fmt.Println(io.EOF) // EOF adalah sebuah error, biasanya digunakan untuk parameter perbandingan
}

// Error yang muncul ketika eksekusi file.Read() akan di-filter, ketika error tersebut adalah selain io.EOF maka proses baca file akan berlanjut. Error io.EOF sendiri menandakan bahwa file yang sedang dibaca adalah baris terakhir isi atau end of file.
