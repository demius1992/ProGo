package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var url = "http://localhost:8090/hello"

//var url = "https://fixer.io/"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	helper := New(RetrySettings{
		BackoffCoef:             2.25,
		IntervalBetweenAttempts: 100 * time.Millisecond,
		MinIntervalAfterFail:    200 * time.Millisecond,
		MaxIntervalAfterFail:    500 * time.Millisecond,
		MaxAttempts:             10,
	}, &Client{})

	if err := helper.Run(ctx, GetInReturn(url)); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("shutting down")
}

func GetInReturn(url string) (job RetryWorker) {
	return func(ctx context.Context) error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(resp.Status)
		fmt.Println(string(body))

		return errors.New(strconv.Itoa(resp.StatusCode))
	}
}
