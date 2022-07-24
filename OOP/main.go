package main

import "fmt"

type Printer interface {
	Print(string)
}

func NewEpson(iInk Ink) Printer { // This is not good because NewEpson should return Epson type not Printer type for best practice
	return &Epson{iInk}
}

type Epson struct {
	ink Ink
}

func (this *Epson) Print(message string) {
	fmt.Println("Printing", message, "With", this.ink.GetColor(), "Color")
}

type Ink struct {
	color string
}

func (this *Ink) GetColor() string {
	return this.color
}

func main() {
	ink := Ink{"Red"}
	epson := NewEpson(ink)
	epson.Print("Hello World")
	epson.(*Epson).ink.GetColor()
}
