package main

import (
	"fmt"
	"net/http"
)

type myError struct {
	Msg        error
	StatusCode int
}

func (m *myError) Error() string {
	return fmt.Sprintf("%s: %d", m.Msg, m.StatusCode)
}

func (m *myError) Is(tgt error) bool {
	fmt.Println("compared!")
	target, ok := tgt.(*myError)
	if !ok {
		return false
	}
	return m.StatusCode == target.StatusCode
}

func ErrorResponse(w http.ResponseWriter, m *myError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.StatusCode)
	resp := fmt.Sprintf(`{"errors":{"message":"%s","type":"%d"}}`, m.Error(), m.StatusCode)
	s := []byte(resp)
	w.Write(s)
}
