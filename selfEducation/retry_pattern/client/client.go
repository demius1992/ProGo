package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ctx     context.Context
	backoff Backoff
}

func NewClient(ctx context.Context, backoff Backoff) *Client {
	return &Client{
		ctx:     ctx,
		backoff: backoff}
}

func main() {
	//url := "https://fixer.io/"
	url := "http://localhost:8090/hello"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	c := NewClient(ctx, Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: false,
	})

	resp, err := c.run(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func (c *Client) run(url string) (string, error) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			resp.Body.Close()
			return "", err
		}

		validate := c.Validate(resp.StatusCode)
		err = c.Retry(validate)
		if err == nil {
			resp.Body.Close()
			return resp.Status, nil
		}
	}
}

func (c *Client) Retry(input interface{}) error {
	check, ok := input.(Action)
	if !ok {
		return errors.New("invalid type casting. Input parameter is not of Action type")
	}

	switch check {
	case Fail:
		return errors.New("fail case, no retrying")
	case Succeed:
		return nil
	case Retry:
		log.Println("trying to reconnect")
		backoff := c.backoff.Duration()
		timeout := time.After(backoff)
		if err := c.sleep(c.ctx, timeout); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) sleep(ctx context.Context, t <-chan time.Time) error {
	select {
	case <-t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
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
