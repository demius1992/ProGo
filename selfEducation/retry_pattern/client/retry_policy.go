package main

type RetryPolicy interface {
	Validate(error) Action
}

type Action int

const (
	Succeed Action = iota
	Fail
	Retry
)
