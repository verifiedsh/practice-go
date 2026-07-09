// Exercise 4: Basic Math API (Multiple Query Parameters)
// Goal: Create a /calculate route that accepts three query parameters: op (operation), a, and b. For example, /calculate?op=add&a=10&b=5 should respond with Result: 15.
// Key Tasks:
// Parse string query variables using r.URL.Query().Get().
// Convert string inputs to integers using strconv.Atoi().
// Support add, subtract, and multiply. Return an HTTP 400 Bad Request if the operation is unknown or parsing fails.
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func mathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path != "/calculate" {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}
		var result int
		op := r.URL.Query().Get("op")
		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")
		x, err := strconv.Atoi(a)
		if err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		switch op {
		case "add":
			result = x + y
		case "subtract":
			result = x - y
		case "multiply":
			result = x * y
		default:
			http.Error(w, "operation unknown", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, result)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
