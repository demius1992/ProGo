package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

type Node struct {
	Next  *Node
	Value interface{}
}

func main() {
	//n := factorial(17)
	//nMinusm := factorial(17 - 6)
	//m := factorial(6)
	//result := n / (m * nMinusm)
	//fmt.Println(real(result))

	start := time.Now()
	writeBuilder("q")
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func writeBuilder(s string) {

	builder := strings.Builder{}
	builder.Grow(1000)

	for i := 0; i < 1000; i++ {
		builder.WriteString(s)
		//fmt.Fprint(&builder, s)
	}
	s1 := builder.String()
	fmt.Println(s1)
}

func writeByteBuffer(s string) {
	buffer := bytes.Buffer{}
	buffer.Grow(1000)
	for i := 0; i < 1000; i++ {
		buffer.WriteString(s)
	}
	s2 := buffer.String()
	fmt.Println(s2)
}

func joining(str string) {
	s := make([]string, 1000)
	count := 0

	for i := 0; i < 1000; i++ {
		s[i] = str

		if count == 10 {
			fmt.Println(strings.Join(s[i-count:i], ""))
			count = 0
		}
		count++
	}
}

//func joining(str string) {
//	s := make([]string, 1000)
//	for i := 0; i < 1000; i++ {
//		s[i] = str
//	}
//	fmt.Println(strings.Join(s, ""))
//}

func factorial(n complex128) complex128 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
