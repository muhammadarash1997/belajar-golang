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
	// myList1 := &List{}
	// myList1.Append(1)
	// myList1.Append(2)
	// myList1.TestingPointer1()

	myList2 := List{}
	myList2.Append(1)
	myList2.Append(2)
	myList2.TestingPointer1()
}

func (this List) TestingPointer1() {
	prev := &this.Head

	fmt.Println(prev)      // 0xc000006028
	fmt.Println(*prev)     // &{1 0xc000088230}
	fmt.Println(this.Head) // &{1 0xc000088230}

	*prev = &Node{}

	fmt.Println(prev)      // 0xc000006028
	fmt.Println(*prev)     // &{0 <nil>}
	fmt.Println(this.Head) // &{0 <nil>}
}

func (this List) TestingPointer2() {
	prev := this.Head

	fmt.Println(prev)      // &{1 0xc000046240}
	fmt.Println(*prev)     // {1 0xc000046240}
	fmt.Println(this.Head) // &{1 0xc000046240}

	*prev = Node{}

	fmt.Println(prev)      // &{0 <nil>}
	fmt.Println(*prev)     // {0 <nil>}
	fmt.Println(this.Head) // &{0 <nil>}
}

func (this List) TestingPointer3() {
	prev := this.Head

	fmt.Println(prev)      // 0xc000006028
	fmt.Println(*prev)     // &{1 0xc000088230}
	fmt.Println(this.Head) // &{1 0xc000088230}

	prev = &Node{}

	fmt.Println(prev)      // 0xc000006028
	fmt.Println(*prev)     // &{0 <nil>}
	fmt.Println(this.Head) // &{0 <nil>}
}