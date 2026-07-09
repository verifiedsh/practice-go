// Exercise 6: Secure Dashboard (Simple Authorization Headers)
// Goal: Create a /dashboard route that acts as a protected page. If the client does not provide a specific API key in their headers, block them.
// Key Tasks:
// Read custom headers using r.Header.Get("X-API-Key").
// Match it against a hardcoded value (e.g., secret123).
// Use http.StatusUnauthorized (401) to reject bad keys.
package main

import (
	"fmt"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/dashboard" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		api := r.Header.Get("X-API-Key")
		if api != "secret123" {
			http.Error(w, "access denied", http.StatusUnauthorized)
			return
		}
		fmt.Fprint(w, "Welcome")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
