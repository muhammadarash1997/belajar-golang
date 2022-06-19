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

func (this *TextManipulator) PrintText() {
	fmt.Println(this.text)
}

func main() {
	myText := NewTextManipulator("Hello World")

	fmt.Println(myText.GetText())
	myText.AppendText("Dude")
	myText.PrintText()
	myText.FindWordAndReplace("Dude", "Guys")
	myText.PrintText()
	myText.FindWordAndDelete("Guys")
	myText.PrintText()
}

// We have to convert this program to be SRP. PrintText method should not be part of TextManipulator
// because it has no cohesion between TextManipulator and PrintText. So therefore we have to separate
// PrintText function by creating new class namely TextPrinter and putting PrintText as a method of the class.