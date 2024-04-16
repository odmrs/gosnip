package main

import (
	"log"
	"net/http"
)

// Define the home handler function which writes a byte slice on body
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from GoSnip"))
}

// Write byte on the new route
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)                         // Put a constant post
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Will use w.WriteHeader(403), and w.Write()
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	// Initialize the server with servermux, will register the home function as the handler for the "/" url pattern
	mux := http.NewServeMux() // Start the serve
	mux.HandleFunc("/", home) // Catch the handler 'home' on the path '/', but will acess all patters

	// To add news routes just put the new patterns inside a handle func
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Using the http.ListenAndServe() function will start a new web server, with 2 parameters
	// tcp network address like ":4000", and the serve mux.
	log.Print("Starting server on :4000")    // LOG show the server was started
	err := http.ListenAndServe(":4000", mux) // Will return error and show the log that error
	log.Fatal(err)
}
