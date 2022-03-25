// BELAJAR POINTER
// NOTE : COBA CEK LANGSUNG KE SETIAP METHOD TestingPointer UNTUK MELIHAT SIFAT-SIFAT DARI POINTER-NYA

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

func main() {
	myList1 := List{}
	myList1.Append(1)
	myList1.Append(2)
	myList1.TestingPointer1()

	fmt.Println("-")
	
	myList2 := &List{}
	myList2.Append(1)
	myList2.Append(2)
	myList2.TestingPointer2()
	
	fmt.Println("-")
	
	myList3 := &List{}
	myList3.Append(1)
	myList3.Append(2)
	myList3.TestingPointer3()

	fmt.Println("-")

	myList4 := &List{}
	myList4.Append(1)
	myList4.Append(2)
	myList4.TestingPointer4()

	fmt.Println("-")
}

func (this *List) TestingPointer1() {
	prev := this.Head
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)

	prev = &Node{}
	// beda ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
	
	this.Head = prev
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
}

func (this *List) TestingPointer2() {
	prev := this.Head
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)

	*prev = Node{}
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
	
	this.Head = prev
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
}

func (this *List) TestingPointer3() {
	prev := this.Head
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)

	*&prev = &Node{}
	// beda ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
	
	this.Head = prev
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
}

func (this *List) TestingPointer4() {
	prev := this.Head
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)

	this.Head = &Node{}
	// beda ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
	
	this.Head = prev
	// sama ketika diprint
	fmt.Println(this.Head)
	fmt.Println(prev)
}