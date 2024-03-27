package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handler)
	fmt.Println("Starting server on Port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}
