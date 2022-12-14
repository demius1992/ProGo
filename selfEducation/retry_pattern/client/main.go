package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

//var url = "http://localhost:8090/hello"
var url = "https://fixer.io/"

type Client struct {
	ctx     context.Context
	backoff Backoff
}

func NewClient(ctx context.Context, backoff Backoff) *Client {
	return &Client{
		ctx:     ctx,
		backoff: backoff,
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := NewClient(ctx, Backoff{Min: 100 * time.Millisecond, Max: 10 * time.Second, Factor: 2, Jitter: false})
	Run(c)

	fmt.Println("shutting down")
}

func Run(retrier Retrier) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		validate := retrier.Validate(resp.StatusCode)
		err = retrier.Retry(validate)

		switch {
		case err == nil:
			fmt.Println(resp.Status)
			return
		case err.Error() == "fail case, no retrying":
			log.Fatalln(err)
		default:
			log.Fatalln(errors.New("timeout is reached"))
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
		return errors.New("transient error, trying to reconnect")
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
