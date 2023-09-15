package main

import (
	"fmt"
	"net/http"
)

// helloHandler responds with "Hello, world!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// movieHandler responds with "Movie info."
func movieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Movie info.")
}

func main() {
	// Register route handlers
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/movie", movieHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
