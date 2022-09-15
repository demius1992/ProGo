package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_retryGet(t *testing.T) {
	expected := "Hello from server"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()
	c := NewClient(svr.URL)

	stop := make(chan string)
	c.retryGet(c.url, stop)

	res := <-stop
	if res != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
