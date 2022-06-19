// INTERFACE DI GOLANG IBARAT ABSTRACT (ABSTRACTION) DI DART

// KESIMPULAN INTERFACE ADALAH SEMUA CLASS/STRUCT YANG TURUNAN DARI INTERFACE WAJIB MEMILIKI METHOD-METHOD YANG ADA DI INTERFACE
// JIKA TIDAK MEMILIKI SEMUA YANG ADA DI INTERFACE MAKA CLASS/STRUCT TSB BUKAN MERUPAKAN TURUNAN DARI INTERFACE TADI
// APABILA TURUNANNYA MEMILIKI METHOD TAMBAHAN SENDIRI MAKA INI TIDAK ERROR SELAGI MEMILIKI SEMUA METHOD-METHODNYA SI INTERFACE, AKAN ERROR JIKA TURUNANNYA TIDAK MEMILIKI SEMUA METHOD YANG ADA DI INTERFACE
// JADI KESIMPULANNYA INTERFACE BERGUNA JIKA INGIN MENGELOMPOKKAN STRUCT-STRUCT YANG MEMILIKI BEHAVIOUR/METHOD-METHOD YANG SAMA

package main

import (
    "fmt"
)

type geometry interface {
    area() int
    perim() int
}

type rectangle struct {
    width int
}
type triangle struct {
    width int
}
type circle struct {    // <-- circle akan juga menjadi turunan interface geometry karena dia turunan triangle yang mana turunan geometry
	triangle
	width int
}

func (r rectangle) area() int {
	return 4
}
func (r rectangle) perim() int {
    return 16
}
func (r rectangle) hello() int {
	return 0
}

func (t triangle) area() int {
    return 3
}
func (t triangle) perim() int {
    return 9
}

func measure(g geometry) {
    fmt.Println(g)
}

func main() {
    r := rectangle{width: 400}
    t := triangle{width: 300}
    c := circle{triangle{width: 700}, 300}

    measure(r)
    measure(t)
    measure(c)
}