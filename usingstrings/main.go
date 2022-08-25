package main

import (
	"fmt"
	"unicode/utf8"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func scanfFunc() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	template := "Product %s %s %f"
	source := "Product Lifejacket Watersports 48.95"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func scanning() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scanln(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
