package main

import "fmt"

type Creature interface {
	Walk()
	Run()
}

type Human struct {
	Creature
}

func NewHuman() Creature {
	return &Human{}
}

func (h *Human) Walk() {
	fmt.Println("Started Walking")
}

func main() {
	h := NewHuman()
	Start(h)
}

func Start(c Creature) {
	c.Walk()
	fmt.Println(c)
}

// SEBELUMNYA YANG KITA TAHU BAHWA SEMUA OBJEK YANG INGIN BERTIPE INTERFACE TERTENTU WAJIB MENGIMPLEMENTASI SEMUA FUNGSI DARI INTERFACE TSB
// TETAPI JIKA KITA MELANGGAR ATURAN TSB DAN OBJEK TSB AKAN TETAP BERTIPE INTERFACE TERTENTU
// DENGAN CARA MENGINJEKSI INTERFACENYA KE DALAM STRUCT OBJEK