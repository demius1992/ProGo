package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 275
	base10String := strconv.FormatInt(int64(val), 10)
	base2String := strconv.FormatInt(int64(val), 2)
	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}
