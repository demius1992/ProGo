package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd := r.Intn(5-1) + 1

	switch rnd {
	case 1:
		ErrorResponse(w, &myError{Msg: errors.New("no success StatusServiceUnavailable"), StatusCode: 503})
		log.Println("No success StatusServiceUnavailable 503")
	case 2:
		ErrorResponse(w, &myError{Msg: errors.New("no success StatusLoopDetected"), StatusCode: 508})
		log.Println("No success StatusLoopDetected 508")
	case 3:
		fmt.Fprint(w, "Hello from server\n")
		log.Println("Success no errors")
	case 4:
		w.WriteHeader(404)
	}

}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
