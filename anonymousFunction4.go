// Anonymous function as an argument
package main

import "fmt"

// Passing anonymous function
// as an argument
func cetak(paramAnony func(p, q string) string) { // nama variable p q tidak mesti sama dengan yg ada di anonymous func yg akan dikirim
	fmt.Println(paramAnony("Geeks", "for"))
}

func main() {
	value := func(a, b string) string { // di sini variabel nya a b bukan p q, ini bisa karena yg dilihat adalah strukturnya, sama juga bisa, malah lebih rapi
		return a + b + "Geeks"
	}

	cetak(value)
}
