package main

import (
	"fmt"
)

func main() {
	d, sieror := divide(5.0, 0.0)
	if sieror != nil {
		fmt.Println(sieror)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Ga bisa dibagi oleh nol boy")
	}
	return a / b, nil
}
