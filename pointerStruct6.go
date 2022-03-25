// METODE PERUBAHAN MELALUI ISI VS METODE PERUBAHAN MELALUI ALAMAT
// PERHATIKAN PADA FUNGSI DELETE PADA BAGIAN --> if current.Data == Data {}

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

func (this *List) PrintList() {
	last := this.Head
	for {
		if last == nil {
			break
		}
		fmt.Println(last.Data)
		last = last.Next
	}
}

func (this *List) Delete(Data int) {
	if this.Head.Data == Data {
		this.Head = this.Head.Next	// perubahan ini menggunakan metode mengubah alamat, bukan menggunakan metode
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
			prev.Next = nil // ini akan mengubah langsung datanya karena .Next dimana kita langsung mengakses datanya, jika ditampung dulu ke variabel misalnya prevNext = prev.Next lalu variabel prevNext kita ubah, maka prev.Next tidak akan berubah
			return
		}

		if current.Data == Data {
			// KITA BISA MENGGUNAKAN METODE 1 ATAU 2 (PILIH SALAH SATU)

			// METODE 1 (PERUBAHAN ISI)
			// metode perubahan di bawah ini menggunakan metode perubahan isi, bukan menggunakan metode perubahan alamat
			*current = *next

			// METODE 2 (PERUBAHAN ALAMAT)
			// metode perubahan di bawah ini menggunakan metode perubahan alamat, bukan menggunakan metode perubahan isi
			/// prev.Next = next
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
	myList := &List{}
	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	myList.Append(4)
	myList.Append(5)
	myList.Append(6)
	myList.PrintList()
	myList.Delete(3)
	myList.PrintList()
}