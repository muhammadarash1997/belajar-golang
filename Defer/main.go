package main

import "fmt"

func main() {
	fmt.Println(Test())
}

func Test() (x int) {
	defer func() {
		fmt.Println("Ini defer", x)
	}()
	return 1
}

// OUTPUT
// Ini defer 1
// 1

// Kesimpulannya defer akan dijalankan setelah return dijalankan,
// karena jika tanpa defer maka hasilnya akan menjadi seperti dibawah ini.

// OUTPUT WITHOUT DEFER
// Ini defer 0
// 1