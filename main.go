// main.go
package main

import (
	"fmt"
	"net/http"
)

// Define a handler function to respond to HTTP requests.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Send a response to the client.
}

func main() {
	// Register the handler function for a specific route ("/").
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
