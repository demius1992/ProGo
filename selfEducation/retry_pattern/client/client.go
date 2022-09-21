package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func main() {
	client := NewClient("http://localhost:8090/hello")
	//client := NewClient("https://fixer.io/")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	retrier := NewRetrier(ctx, &client, Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: false,
	})

	responce, err := retrier.Run(client.url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(responce)
}

func (c *Client) Validate(resp interface{}) Action {
	res, ok := resp.(int)
	if !ok {
		log.Println("invalid status code")
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
