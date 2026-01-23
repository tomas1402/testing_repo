package main

import (
	"fmt"
	"net/http"
	"hello-project/internal/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)
	
	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}