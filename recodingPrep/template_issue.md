# template.Template usage in go server to parse html files in templates folder

To parse HTML files from a templates folder in a Go server, always use the html/template package instead of text/template to ensure contextual auto-escaping against cross-site scripting (XSS) attacks.

The most efficient way to load an entire directory of templates is by utilizing template.ParseGlob or Go's embed filesystem.

# Project Directory Structure
text.

├── main.go
└── templates/
    ├── index.html
    └── about.html


# Method 1: Using template.ParseGlob (Recommended for Local Dev)

This method reads the files directly from your disk at runtime. It matches a wildcard pattern to load all files inside the folder.


```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Global variable to hold parsed templates
var tmpl *template.Template

func init() {
	// ParseGlob parses all files matching the pattern inside the templates folder
	// template.Must panics if parsing fails (e.g., syntax error or missing directory)
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"Title": "Home Page", "Message": "Welcome to Go!"}
	
	// ExecuteTemplate targets a specific file name within the parsed bundle
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```













**what happens if you do not declare var tmpl *template.Template  and just declare var tmpl = template.Must(template.ParseGlob("templates/*.html")) outside the function body**

If you use 
var tmpl = template.Must(template.ParseGlob("templates/*.html")) 
at the package level (outside any function body), your code will compile and work perfectly.

This is standard, idiomatic Go for global variable declaration. Go automatically infers that tmpl is of type *template.Template because that is what template.Must returns.

However, doing this introduces one critical runtime behavior you must watch out for: Initialization Timing.

**The Initialization Order Trap**

In Go, package-level variables are initialized before the init() function runs and before the main() function executes.

This means template.ParseGlob("templates/*.html") will run the exact millisecond your program starts up.

The Problem: 
If your server runs in an environment where the working directory changes, or if you run your tests (go test ./...) from a different subfolder, the application will look for the templates/ folder relative to where the command was run.

The Result: 
If it cannot find the folder instantly at startup, template.Must will trigger an immediate runtime panic and crash your app before main() even starts.

**Compare the Two Valid Approaches**

Approach                 |Code                                                  |Pros & Cons                                                                                |
-------------------------|------------------------------------------------------|-------------------------------------------------------------------------------------------|
Inferred Global          |var tmpl = template.Must(...)                         |🟢 Clean, short code.                                                                      |
(Your suggestion)        |                                                      |🔴 Panics immediately at startup if paths are slightly off (e.g., during testing).         |
|
Explicit Global +        |var tmpl *template.Templatefunc init() { tmpl = ... } |🟢 Allows you to log custom errors or dynamically adjust the directory path before parsing.| 
init()                   |                                                      |🔴 Slightly more boilerplate code.                                                         |


**How to make your approach 100% safe**
If you want to use the short var tmpl = ... syntax outside your functions safely, pair it with go:embed (as shown in Method 2 of the first response).go


```go
package main

import (
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templateFiles embed.FS

// 🟢 This is 100% safe because the files are baked into the binary at compile time.
// It will never panic at runtime due to missing local folders.
var tmpl = template.Must(template.ParseFS(templateFiles, "templates/*.html"))

func main() {
    // Your server code here
}
```