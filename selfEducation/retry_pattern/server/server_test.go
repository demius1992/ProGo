package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_hello(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	hello(w, request)
	resp := w.Result()
	defer resp.Body.Close()

	headerContentType := resp.Header.Get("Content-Type")
	expectedEmpty := ""
	expectedJson := "application/json"
	expected := "text/plain; charset=utf-8"

	if headerContentType != expectedJson && headerContentType != expectedEmpty && headerContentType != expected {
		t.Errorf("incorrect content type: got %v", headerContentType)
	}

	headerStatusCode := resp.StatusCode
	expectedOK := http.StatusOK
	expectedSSU := http.StatusServiceUnavailable
	expectedLoop := http.StatusLoopDetected
	expectedNotFound := http.StatusNotFound

	if headerStatusCode != expectedOK && headerStatusCode != expectedSSU && headerStatusCode != expectedLoop && headerStatusCode != expectedNotFound {
		t.Errorf("handler returned wrong status code: got %v", headerStatusCode)
	}
}
