// Exercise 3: Text Counter (URL Variables & Methods)
// Goal: Build a server with a /count route. If a user sends a GET request, return the text "Send a POST request with text to count words". If they send a POST request, read the text body and return the number of characters.
// Key Tasks:
// Differentiate between GET and POST methods using r.Method.
// Read the entire request body using io.ReadAll(r.Body).
// Return the character length as a string.
package main

import (
	"fmt"
	"io"
	"net/http"
)

func counterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/count" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, "Send a POST request with text to count words")
	case "POST":
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprint(w, len(data))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
