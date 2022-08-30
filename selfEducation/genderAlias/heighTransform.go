package main

import "fmt"

type Gender string

type man struct {
	gender Gender
	height float64
}

type woman struct {
	gender Gender
	height float64
}

func (g Gender) heightTransform(sm int) float64 {
	return float64(sm) / 30.48
}

func main() {
	m := man{gender: "male"}
	m.height = m.gender.heightTransform(170)

	fmt.Printf("height in foot is %.2f \n", m.height)
	fmt.Println("gender is " + m.gender)
}
