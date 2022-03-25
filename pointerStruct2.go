// NOTE : PERHATIKAN FUNGSI/METHOD DELETE, KARENA MATERI POINTER INI BERADA DI DALAM FUNGSI/METHOD DELETE DAN AMATI ANTARA DELETE1 DAN DELETE2

package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func (this *List) Append(Data int) {
	newNode := &Node{Data: Data}

	if this.Head == nil {
		this.Head = newNode
		return
	}

	thisNode := this.Head
	for {
		if thisNode.Next == nil {
			thisNode.Next = newNode
			break
		}
		thisNode = thisNode.Next
	}
}

func (this *List) PrintList() {
	fmt.Println("PrintList dipanggil")
	last := this.Head
	for {
		if last == nil {
			break
		}
		fmt.Println(last.Data)
		last = last.Next
	}
}

func main() {
	myList := List{}
	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	myList.Append(4)

	// DIPANGGIL PERTAMA KALI
	myList.PrintList()

	fmt.Println("-")

	myList.Delete()

	fmt.Println("-")

	// DIPANGGIL KEDUA KALI
	myList.PrintList()
}

func (this *List) Delete() {
	prev := &this.Head
	current := &(*prev).Next
	next := &(*current).Next

	fmt.Println("nilai prev", (*prev).Data)
	fmt.Println("nilai current", (*current).Data)
	fmt.Println("nilai next", (*next).Data)

	// COBA SATU-SATU METODE 1 DAN 2 DALAM MENGUBAH NILAI, HASIL PrintList AKAN BERBEDA KETIKA DIPANGGIL KEDUA KALINYA JIKA METODE 1 DAN 2 DIBANDINGKAN

	// METODE 1 PADA METODE INI TIDAK AKAN MENGUBAH NILAI ASLINYA

	prev = current       // <-- Metode ini tidak akan mengubah nilai asli karena prev, current dan next hanya akan menyimpan alamat baru
	current = next       // Variabel current diisi dengan alamat baru
	next = &(*next).Next // <-- Coba next 1, Variabel next diisi dengan alamat baru
	// next = nil        // <-- Coba next 2, Variabel next diisi dengan nil, jadi next tidak akan lagi menyimpan alamat yang sebelumnya disimpan olehnya
	// ^ DIATAS ADA 2 NEXT, COBALAH SATU PERSATU

	// METODE 2 PADA METODE INI AKAN MENGUBAH NILAI ASLINYA

	// *prev = *current     // <-- Metode ini akan mengubah nilai asli karena isi dari alamat yang disimpan oleh prev, current dan next diisi nilai baru bukan alamat, sehingga nilai asli berubah begitu juga nilai prev, current dan next
	// *current = *next     // Mengubah isi dari alamat yang disimpan oleh current
	// *next = (*next).Next // <-- Coba next 1, Mengubah isi dari alamat yang disimpan oleh next
	// *next = nil			// <-- Coba next 2, Mengubah isi dari alamat yang disimpan oleh next, jadi nilai aslinya juga akan berubah jadi nil
	// ^ DIATAS ADA 2 NEXT, COBALAH SATU PERSATU

	fmt.Println("nilai prev", (*prev).Data)
	fmt.Println("nilai current", (*current).Data)
	fmt.Println("nilai next", (*next).Data) // <-- Ketika mencoba next 2, maka ini fungsi pada baris ini harus dinon-aktifkan, jika tidak akan error, karena ketika sudah diisi nil otomatis next tidak akan mempunyai field Data
}
