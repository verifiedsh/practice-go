package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/hello", queryHandler)
	http.HandleFunc("/count", counterHandler)
	http.HandleFunc("/calculate", mathHandler)
	http.HandleFunc("/agent", agentHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/legacy", redirectorHandler)
	http.HandleFunc("/v2", redirectedHandler)
	fmt.Println("Server listening live @:8080...")
	http.ListenAndServe(":8080", nil)
}

// The use of the internet or external resources is not allowed.
// https://rebrand.ly/lettherebecode
// 1- Repo to submit when you are done: http-haven
// 2- Submit each exercise as a separate file in the http-haven directory.
// 3- Collaboration with peers is permitted. Your fellows are the ONLY source of truth.
// Terminal Test Script (test_endpoints.sh)
// The shell test script acts as an automated collection of curl requests. You can run this directly from their terminal to verify if their 7 endpoints match your performance requirements.
// Create a file named test_endpoints.sh, paste the code below, and make it executable using chmod +x test_endpoints.sh.
// Use ./test_endpoints.sh  to run the tests.

/*
Go HTTP Server Cheat Sheet
******Core Standard Library Functions******

1. 		Routing & Server Management (net/http)
http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))
Registers a handler function for a specific URL path pattern.
***    	http.ListenAndServe(addr string, handler Handler) error
Starts an HTTP server on the specified address (e.g., ":8080"). Passing nil uses the default system router (DefaultServeMux).
***		http.Error(w ResponseWriter, error string, code int)
A helper that sends a specific error message string and numeric HTTP status code back to the client, automatically ending the request lifecycle safely.
***		http.Redirect(w ResponseWriter, r *Request, url string, code int)
Sends an HTTP redirect status code (like 301 or 302) forcing the client's browser to jump to a new target URL path.

2. 		Reading Inputs & Formatting (io, strconv, fmt)
***		io.ReadAll(r io.Reader) ([]byte, error)
Reads all remaining data from an input stream (like r.Body) until it hits the end of the file/stream (EOF). Returns a byte slice.
***		strconv.Atoi(s string) (int, error)
Stands for "ASCII to Integer". Converts a text string into a native numeric int. Returns an error if the text contains non-numeric characters.
***		mt.Fprintf(w io.Writer, format string, a ...any)
Formats data according to a template string and writes the output directly into an open network socket stream or file (w).



Request Context Breakdown (*http.Request)
When a client hits your server, all incoming information is bundled inside the pointer to the http.Request struct (usually named r):
Struct Field / Method
Purpose / Explanation
Example Usage
r.Method
A string representing the incoming HTTP type. Always use standard library constants for comparisons.
if r.Method == http.MethodPost
r.Body
An open stream containing data uploaded via POST/PUT requests. Always close it after reading to prevent memory leaks.
defer r.Body.Close()
r.URL.Query()
Parses the raw URL query string (everything after the ?) and returns a map-like structure (Values).
queryMap := r.URL.Query()
r.URL.Query().Get(key)
Fetches the value of a specific query parameter. Returns an empty string "" if the key does not exist.
name := r.URL.Query().Get("name")
r.Header.Get(key)
Fetches the value of a specific HTTP request header (case-insensitive). Returns "" if missing.
token := r.Header.Get("X-API-Key")





Response Management (http.ResponseWriter)
The http.ResponseWriter interface (usually named w) is your pipeline to send data back to the client. Keep this execution order in mind:
Modify Headers First: Call w.Header().Set("Key", "Value") before anything else if changing content types or metadata.
Write Status Code Second: Call w.WriteHeader(http.StatusCreated) if returning something other than a standard 200 OK.
Write Body Content Last: Call w.Write([]byte) or fmt.Fprintf(w) to push actual visual content to the user. Writing to the body locks your headers and status code automatically.



*/
