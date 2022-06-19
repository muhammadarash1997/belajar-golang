// PERHATIKAN PADA FUNGSI LENGTH
// PERHATIKAN PERUBAHAN-PERUBAHANNYA

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

func main() {
	myList := List{}
	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	myList.Append(4)
	myList.Length()
}

func (this *List) Length() {
	prev := this.Head
	current := prev.Next
	next := current.Next

	fmt.Println("--------------------------------------------------[1]")
	fmt.Println("thisHead", this.Head)

	fmt.Println("--------------------------------------------------[2]")
	fmt.Println("prev before", prev)
	fmt.Println("current before", current)
	fmt.Println("next before", next)

	fmt.Println("--------------------------------------------------[3]")
	fmt.Println("prev.Next before", prev.Next)
	fmt.Println("current.Next before", current.Next)
	fmt.Println("next.Next before", next.Next)
	
	fmt.Println("--------------------------------------------------[4]")
	fmt.Println("*prev.Next", *prev.Next)
	fmt.Println("*current.Next before", *current.Next)
	fmt.Println("*next.Next before", *next.Next)
	
	*current = *prev

	fmt.Println("--------------------------------------------------[5]")
	fmt.Println("prev after", prev)
	fmt.Println("current after", current)
	fmt.Println("next after", next)
	
	fmt.Println("--------------------------------------------------[6]")
	fmt.Println("prev.Next after", prev.Next)
	fmt.Println("current.Next after", current.Next)
	fmt.Println("next.Next after", next.Next)
	
	fmt.Println("--------------------------------------------------[7]")
	fmt.Println(this.Head)
}
