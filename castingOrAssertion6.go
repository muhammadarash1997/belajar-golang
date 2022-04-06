package main

import "io"

// BAD
type A struct{}

func New1() A {
	return A{}
}
func (a A) Read(p []byte) (n int, err error) {
	panic("implement me")
}
func (a A) Close() error {
	panic("implement me")
}

// GOOD
type B struct{}

var _ io.ReadCloser = (*B)(nil)	// <-- Dengan menggunakan ini kita jadi tahu bahwa sebuah struct mengimplementasikan sebuah interface

func New2() B {
	return B{}
}

func (b B) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (b B) Close() error {
	panic("implement me")
}
