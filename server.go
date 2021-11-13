package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configure path and handler function
	http.HandleFunc("/hello", Hello)
	// Listen on port 8080 and block main
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
