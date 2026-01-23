package handlers

import (
	"net/http"
)

// Hello is a simple handler that returns a greeting
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}