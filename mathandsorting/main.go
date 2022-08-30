package main

import (
	"fmt"
	"github.com/rs/xid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//myGenerator(18)
	//generator()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func generator() {
	guid := xid.New()
	println(guid.String())
}

func myGenerator(length int) {
	var builder strings.Builder
	resultSlice := make([]string, length)

	for i := 0; i < length/3; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		intValue := strconv.Itoa(r.Intn(10-1) + 1)
		charUp := r.Intn(90-65) + 65
		charLow := r.Intn(122-97) + 97

		intFieldsUp := string(rune(charUp))
		intFieldsDown := string(rune(charLow))

		resultSlice = append(resultSlice, intValue, intFieldsUp, intFieldsDown)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(resultSlice), func(i, j int) { resultSlice[i], resultSlice[j] = resultSlice[j], resultSlice[i] })
	for _, a := range resultSlice {
		builder.WriteString(a)
	}
	resultString := builder.String()

	fmt.Println(resultString)
}
