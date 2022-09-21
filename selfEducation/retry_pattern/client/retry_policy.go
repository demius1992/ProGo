package main

type RetryPolicy interface {
	Validate(interface{}) Action
}

type Action int

const (
	Succeed Action = iota
	Fail
	Retry
)
