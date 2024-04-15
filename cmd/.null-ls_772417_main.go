package main

import (
	"log"
	"net/http"
)

// Define the home handler function which wirtes a byte slice on body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from GoSnip"))
}

func main() {
	// Initialize the server with servermux, will register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux() // Start the serve
	mux.HandleFunc("/", home) // Catch the handler 'home' on the path '/'

	// Using the http.ListenAndServe() function will start a new web server, with 2 parameters
	// tcp network address like ":4000", and the serve mux.

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
