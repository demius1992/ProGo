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
	start := time.Now()
	writeByteBuffer("q")
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func writeBuilder(s string) {
	builder := strings.Builder{}
	for i := 0; i < 1000; i++ {
		builder.WriteString(s)
	}
	s1 := builder.String()
	fmt.Println(s1)
}

func writeByteBuffer(s string) {
	buffer := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		buffer.WriteString(s)
	}
	s2 := buffer.String()
	fmt.Println(s2)
}

func joining(str string) {
	s := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s = append(s, str)
	}
	fmt.Println(strings.Join(s, ""))
}
