package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Retrier struct {
	ctx         context.Context
	retryPolicy RetryPolicy
	backoff     Backoff
}

func NewRetrier(ctx context.Context, retryPolicy RetryPolicy, backoff Backoff) Retrier {
	return Retrier{
		ctx:         ctx,
		retryPolicy: retryPolicy,
		backoff:     backoff,
	}
}

func (r *Retrier) Run(url string) (interface{}, error) {
	for {

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		switch r.retryPolicy.Validate(resp.StatusCode) {
		case Fail:
			return "Failed", err
		case Succeed:
			return resp.Status, nil
		case Retry:
			log.Println("trying to reconnect")

			backoff := r.backoff.Duration()
			timeout := time.After(backoff)
			if err = r.sleep(r.ctx, timeout); err != nil {
				return nil, err
			}
		}
	}
}

func (r *Retrier) sleep(ctx context.Context, t <-chan time.Time) error {
	select {
	case <-t:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
