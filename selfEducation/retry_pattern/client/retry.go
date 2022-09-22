package main

type Retrier interface {
	Retry(interface{}) error
}
