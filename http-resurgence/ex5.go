// Exercise 5: Status Code Factory
// Goal
// Build a /status endpoint that accepts a code query parameter containing any HTTP status code number. The server must respond using that exact status code. This forces you to think about how status codes are set — and when they cannot be changed.

// Key Tasks
// ●     Read the code query parameter using r.URL.Query().Get("code").
// ●     If code is missing or empty — return 400 with the message: "code parameter is required".
// ●     Convert code to an integer using strconv.Atoi(). If conversion fails — return 400 with: "code must be a valid integer".
// ●     If the integer is not between 100 and 599 — return 400 with: "code must be a valid HTTP status code (100–599)".
// ●     If the code is valid — respond with that exact status code using w.WriteHeader(code) and a body message: "Responding with status [code]".

// Critical rule —
// You must call w.WriteHeader(code) BEFORE writing anything to w with w.Write()
// or fmt.Fprintf(). If you write the body first, Go automatically sends a 200 header
// and you cannot change it afterwards. The order is: WriteHeader → then Write.
// Test this deliberately: call w.Write() first, then w.WriteHeader(404). What does
// curl -v show you? Write your observation in a comment in your file.

// Stretch — do this after the core task works
// ●     After calling w.WriteHeader(), append a descriptive name to the body message.
// ○     ?code=404 → "Responding with status 404 Not Found"
// ○     Use http.StatusText(code) to get the official status name.
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/status" {
	// 	http.Error(w, "page not found", http.StatusNotFound)
	// 	return
	// }
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	statusCode, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}
	if statusCode < 100 || statusCode > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", http.StatusBadRequest)
		return
	}
	// w.Write([]byte(code))
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Responding with status %d", statusCode)
	// call w.Write() first, then w.WriteHeader(404) results in
}
