package main

import "fmt"

type taxSystem interface {
    calculateTax() int
}

type indianTax struct {
    taxPercentage int
    income        int
}

func (i *indianTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}

type singaporeTax struct {
    taxPercentage int
    income        int
}

func (i *singaporeTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}

func main() {
    indianTax := &indianTax{
        taxPercentage: 30,
        income:        1000,
    }
    singaporeTax := &singaporeTax{
        taxPercentage: 10,
        income:        2000,
    }


    taxSystems := []taxSystem{indianTax, singaporeTax}	// <-- Di sini terlihat keuntungan dari menggunakan interface, jadi lebih seperti turunan
    totalTax := calculateTotalTax(taxSystems)


    fmt.Printf("Total Tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
    totalTax := 0
    for _, t := range taxSystems {
        totalTax += t.calculateTax() //This is where runtime polymorphism happens
    }
    return totalTax
}

// Interface. Jika suatu struct ingin mengimplementasi interface tertentu maka
// struct tersebut wajib memiliki semua fungsi dari interface tsb, tidak bisa
// hanya salah satu fungsi saja.
// Aturan ini berlaku di semua bahasa. Jika suatu class mengimplementasikan
// suatu interface maka class tsb wajib memiliki semua function yang ada di
// interface tsb.