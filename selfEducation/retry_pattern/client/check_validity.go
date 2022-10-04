package main

import (
	"strconv"
)

type Client struct{}

func (c *Client) Validate(err error) Action {
	if err == nil {
		return Fail
	}

	res, err := strconv.Atoi(err.Error())
	if err != nil {
		return Fail
	}

	switch {
	case res >= 200 && res <= 226:
		return Succeed
	case res >= 400 && res <= 451:
		return Fail
	default:
		return Retry
	}
}
