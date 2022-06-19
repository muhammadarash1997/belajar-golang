// POINTER
// PERHATIKAN LANGUSNG KE BAGIAN METHOD DELETE
// ADA DUA METHOD DELETE
// YANG PERTAMA MENGGUNAKAN **NODE <-- YANG MANA MENYIMPAN ALAMAT DARI YANG MENYIMPAN ALAMAT SI NODE NYA, CARA PENGAKSESANNYA (*current).Data, LIHAT CARA PENGAKSESANNYA BERBEDA DENGAN YANG KEDUA
// YANG KEDUA MENGGUNAKAN *NODE <-- YANG MANA MENYIMPAN ALAMAT SI NODE NYA LANGSUNG, CARA PENGAKSESANNYA current.Data, LIHAT CARA PENGAKSESANNYA BERBEDA DENGAN YANG PERTAMA
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
			return
		}
		thisNode = thisNode.Next
	}
}

// FUNGSI DELETE YANG INI MENGGUNAKAN VARIABEL **NODE
func (this *List) Delete1(Data int) {
	if this.Head.Data == Data {
		this.Head = this.Head.Next
		return
	}

	prev := &this.Head
	current := &(*prev).Next
	next := &(*current).Next

	for {
		if next == nil && (*current).Data != Data {
			fmt.Println("Data not-found")
			return
		}

		if next == nil && (*current).Data == Data {
			(*prev).Next = nil
			return
		}

		if (*current).Data == Data {
			(*prev).Next = *next
			return
		}

		prev = current
		current = next

		if (*current).Next == nil {
			next = nil
			continue
		}

		next = &(*current).Next
	}
}

// FUNGSI DELETE YANG INI MENGGUNAKAN VARIABEL *NODE
func (this *List) Delete2(Data int) {
	if this.Head.Data == Data {
		this.Head = this.Head.Next
		return
	}

	prev := this.Head
	current := prev.Next
	next := current.Next

	for {
		if next == nil && current.Data != Data {
			fmt.Println("Data not-found")
			return
		}

		if next == nil && current.Data == Data {
			prev.Next = nil
			return
		}

		if current.Data == Data {
			prev.Next = next
			return
		}

		prev = current
		current = next

		if current.Next == nil {
			next = nil
			continue
		}
		next = current.Next
	}
}

func main() {
	myList := List{}
	myList.Append(1)
	myList.Append(2)
	myList.Delete1(2)
	myList.Delete2(2)
}
