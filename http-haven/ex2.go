// Exercise 2: Query Parameters & Path Validation
// Goal: Create a /hello endpoint that reads a name query parameter (e.g., /hello?name=Alice) and responds with "Hello, Alice!". If the parameter is missing, default to "Hello, Guest!".
// Tasks:
// Extract query parameters using r.URL.Query().Get("name").
// Reject any HTTP method that is not GET by returning an http.StatusMethodNotAllowed status code.
package main

import (
	"fmt"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/hello" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest!"
		}
		fmt.Fprintf(w, "Hello, %v!", name)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
