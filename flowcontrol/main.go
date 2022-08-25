package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 := &p1
	p1.name = "Original Kayak"
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", (*p2).name)
}
