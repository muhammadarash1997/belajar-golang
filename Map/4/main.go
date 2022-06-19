package main

import (
	"fmt"
)

type node struct {
	number  int
	faction string
}

func (n *node) setNumber(number int) {
	n.number = number
}

func main() {
	k := make(map[string]interface{})
	a := node{}
	a.setNumber(1)
	k["aaa"] = &a

	i := node{}
	i.setNumber(1)
	k["iii"] = &i

	u := node{}
	u.setNumber(3)
	k["uuu"] = &u

	e := node{}
	e.setNumber(4)
	k["eee"] = &e

	var count int
	for _, d := range k {
		fmt.Println(d)
		count++
		if count == 4 {

			o := node{}
			o.setNumber(5)
			k["ooo"] = o

			for _, e := range k {
				fmt.Println(e)
			}
		}
	}
}
