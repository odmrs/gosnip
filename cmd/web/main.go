package main

import (
	"log"
	"net/http"
	// Import the package handlers to use different functions inside there package
	handlers "github.com/odmrs/gosnip/internal"
)

func main() {
	// Initialize the server with servermux, will register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux()          // Start the serve
	mux.HandleFunc("/", handlers.Home) // Catch the handler 'home' on the path '/', but will acess all patters
	// To add news routes just put the new patterns inside a handle func
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	// Using the http.ListenAndServe() function will start a new web server, with 2 parameters
	// tcp network address like ":4000", and the serve mux.
	log.Print("Starting server on :4000")    // LOG show the server was started
	err := http.ListenAndServe(":4000", mux) // Will return error and show the log that error
	log.Fatal(err)
}
