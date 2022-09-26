package main

import (
	"errors"
	"fmt"
)

type Gender string

const (
	male   Gender = "male"
	female Gender = "female"
)

func (g Gender) heightTransform(sm int) float64 {
	return float64(sm) / 30.48
}

func (g Gender) parseGender(str string) (Gender, error) {
	if string(g) == str {
		return Gender(str), nil
	}
	return "", &customError{Msg: "wrong gender", Gender: string(g)}
}

func main() {

	gender, err := female.parseGender("male")

	var custom = &customError{
		Msg:    "wrong gender",
		Gender: "female",
	}

	if err != nil {
		if errors.Is(err, custom) {
			fmt.Printf("parsing gender is failed because of %s, expected is %s", custom.Msg, custom.Gender)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("gender is %s", gender)
}
