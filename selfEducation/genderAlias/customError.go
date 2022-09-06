package main

import (
	"fmt"
)

type customError struct {
	Msg    string
	Gender string
}

func (c *customError) Error() string {
	return fmt.Sprintf("%s: %s", c.Msg, c.Gender)
}

func (c *customError) Is(tgt error) bool {
	fmt.Println("compared!")
	target, ok := tgt.(*customError)
	if !ok {
		return false
	}
	return c.Gender == target.Gender
}
