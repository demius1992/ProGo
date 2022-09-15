package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
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
	stop := make(chan string, 1)

	client.retryGet(client.url, stop)

	responce := <-stop

	fmt.Println(responce)
}

func (c Client) retryGet(url string, stop chan string) {
	counter := 1

	go func() {
		for {
		start:
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln("error is not nil", err)
			}

			if resp.StatusCode == http.StatusServiceUnavailable {
				fmt.Printf("Failed to connect server, no success StatusServiceUnavailable. Responce status: %d, here we go again\n",
					resp.StatusCode)

				time.Sleep(time.Second * time.Duration(counter))
				counter++
				goto start

			} else if resp.StatusCode == http.StatusLoopDetected {
				fmt.Printf("Failed to connect server, no success StatusLoopDetected. Responce status: %d, here we go again\n",
					resp.StatusCode)

				time.Sleep(time.Second * time.Duration(counter))
				counter++
				goto start

			} else if resp.StatusCode == http.StatusOK {
				fmt.Println("Response status:", resp.Status)

				out, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(errors.Wrap(err, "unable to read response data"))
				}

				defer resp.Body.Close()
				stop <- string(out)
				return

			} else {
				fmt.Printf("unexpected behavior: %s, let's try to reconnect\n", resp.Status)
				time.Sleep(time.Second * time.Duration(counter))
				counter++
				goto start
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)
		stop <- "request time out"
	}()
}
