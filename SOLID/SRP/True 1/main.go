package main

import (
	"fmt"
	"strings"
)

type TextManipulator struct {
	text string
}

func NewTextManipulator(text string) *TextManipulator {
	return &TextManipulator{text}
}

func (this *TextManipulator) GetText() string {
	return this.text
}

func (this *TextManipulator) AppendText(newText string) {
	this.text = this.text + " " + newText
}

func (this *TextManipulator) FindWordAndReplace(word string, replacementWord string) string {
	if strings.Contains(this.text, word) {
		this.text = strings.Replace(this.text, word, replacementWord, -1)
	}

	return this.text
}

func (this *TextManipulator) FindWordAndDelete(word string) string {
	if strings.Contains(this.text, word) {
		this.text = strings.Replace(this.text, word, "", -1)
	}
	return this.text
}

type TextPrinter struct {
	textManipulator *TextManipulator
}

func NewTextPrinter(textManipulator *TextManipulator) TextPrinter {
	return TextPrinter{textManipulator}
}

func (this *TextPrinter) PrintText() {
	fmt.Println(this.textManipulator.GetText())
}

func (this *TextPrinter) PrintOutEachWordOfText() {
	fmt.Println(strings.Split(this.textManipulator.text, " "))
}

func main() {
	myTextManipulator := NewTextManipulator("Hello World")
	myTextPrinter := NewTextPrinter(myTextManipulator)

	fmt.Println(myTextManipulator.GetText())
	myTextManipulator.AppendText("Dude")
	myTextPrinter.PrintText()
	myTextPrinter.PrintOutEachWordOfText()

	myTextManipulator.FindWordAndReplace("Dude", "Guys")
	myTextPrinter.PrintText()
	myTextPrinter.PrintOutEachWordOfText()
	
	myTextManipulator.FindWordAndDelete("Guys")
	myTextPrinter.PrintText()
	myTextPrinter.PrintOutEachWordOfText()
}