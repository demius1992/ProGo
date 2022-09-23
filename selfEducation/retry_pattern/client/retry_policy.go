package main

type Retrier interface {
	Validate(interface{}) Action
	Retry(interface{}) error
}

type Action int

const (
	Succeed Action = iota
	Fail
	Retry
)
