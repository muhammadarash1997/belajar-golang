package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var amount = []int{5, 6, 7, 8}

func randomFirstName() string {
	rand.Seed(time.Now().UnixNano())
	length := 3 + rand.Intn(5)
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomLastName() string {
	rand.Seed(time.Now().UnixNano())
	length := 3 + rand.Intn(5)
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestRandomGenerator(t *testing.T) {
	fmt.Println(randomFirstName())
	fmt.Println(randomLastName())
}
