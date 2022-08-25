package main

import (
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func index[T comparable](s []T, x T) int {
	for i, t := range s {
		if t == x {
			return i
		}
	}
	return -1
}

func main() {

	si := []int{1, 2, 3, 4, 5}
	ss := []string{"hello", "qwerty", "foo"}

	fmt.Println(index(si, 4))
	fmt.Println(index(ss, "qwerty"))
}
