// Exercise 4: Form Decoder
// Goal
// Build a /form endpoint that accepts a POST request with a URL-encoded form body containing two fields: username and language. Parse both, validate them, and return a formatted confirmation. This is the closest exercise in this set to what ascii-art-web actually does.

// Key Tasks
// ●     Reject non-POST requests with 405.
// ●     Call r.ParseForm() to parse the incoming form body — do not skip this step.
// ●     Read username and language using r.FormValue().
// ●     If either field is empty — return 400 Bad Request with a message identifying which field is missing.
// ○     Missing username → "username is required"
// ○     Missing language → "language is required"
// ●     If both are present — respond with: "Hello [username], you are coding in [language]!"

// Think about —
// What is the difference between r.ParseForm() + r.Form.Get() and just r.FormValue()?
// r.FormValue() calls ParseForm internally — but calling ParseForm explicitly first
// gives you control over error handling. Try it both ways.

// Stretch — do this after the core task works
// ●     Handle the case where the request Content-Type is not application/x-www-form-urlencoded. Return a 415 Unsupported Media Type with a clear message.
// ○     Read Content-Type from r.Header.Get() and check it before parsing.
// ○     Test it: curl -X POST -H "Content-Type: text/plain" -d "username=Ada" http://localhost:8080/form
package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/form" {
	// 	http.Error(w, "page not found", http.StatusNotFound)
	// 	return
	// }
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			http.Error(w, "invalid content type", http.StatusUnsupportedMediaType)
			return
		}
		username := r.FormValue("username")

		language := r.FormValue("language")
		if username == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}
		if language == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %v, you are coding in %v!", username, language)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
