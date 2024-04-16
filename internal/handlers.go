package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define the home handler function which writes a byte slice on body
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from GoSnip"))
}

// Write byte on the new route
func SnippetView(w http.ResponseWriter, r *http.Request) {
	// Will get the id parameter from query, if return err or the id is less than 1, we return 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)                         // Put a constant post
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Will use w.WriteHeader(403), and w.Write()
		return
	}

	w.Write([]byte("Create a new snippet"))
}
