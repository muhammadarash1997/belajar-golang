package main

import "fmt"

type regions struct {
	region map[string]int
}

func (r *regions) append(d string) {
	if r.region == nil {
		r.region = make(map[string]int)
	}

	r.region[d]++
}

func main() {
	a := regions{}

	a.append("d")
	a.append("a")
	a.append("d")

	fmt.Println(a)
}
