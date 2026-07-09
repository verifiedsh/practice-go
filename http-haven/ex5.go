// Exercise 5: User-Agent Echo (Reading Headers)
// Goal: Create an /agent route that reads the incoming browser or client header details and echoes it back in plain text: "You are visiting us using: [User-Agent Info]".
// Key Tasks:
// Inspect request headers using r.Header.Get("User-Agent").
// Handle instances where the header might be missing or empty.

package main

import (
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/agent" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		agent := r.Header.Get("User-Agent")
		if agent == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "You are visiting us using: %v", agent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
