package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC()) // Akan sama dengan waktu pusat dunia
	fmt.Println(time.Now())
}
