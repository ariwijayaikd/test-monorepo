package main

import (
	"fmt"
	"net/http"
)

func GoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! This is a test response from the Go server.")
}

func main() {
	http.HandleFunc("/", GoHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
