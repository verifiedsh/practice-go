# Exercise 1 — Custom Error Pages

## Goal

Replace every call to `http.Error()` in your application with reusable HTML error pages.

Instead of returning plain text errors like:

```go
http.Error(w, "404 Not Found", http.StatusNotFound)
```

your server should render the appropriate template while still sending the correct HTTP status code.

---

## Background

Browsers understand much more than plain text—they can display HTML pages for errors.

Instead of sending:

```
404 Not Found
```

you'll create pages such as:

```
templates/
    404.html
    405.html
    500.html
```

Since multiple handlers may return the same errors, repeating the template-rendering code everywhere violates the DRY (Don't Repeat Yourself) principle.

The goal is to centralize error handling into a helper function.

---

## Key Tasks

* Create separate templates for:

  * 404 Not Found
  * 405 Method Not Allowed
  * 500 Internal Server Error

* Create a reusable helper function that:

  * receives a `ResponseWriter`
  * receives the template name
  * receives the status code
  * writes the correct status code
  * renders the HTML page

* Replace every use of `http.Error()` throughout your project with this helper.

* Ensure the browser receives both:

  * the correct HTML page
  * the correct HTTP status code

---

## Success Criteria

These URLs should behave correctly:

```
GET /does-not-exist
```

Displays your custom 404 page.

```
POST /
```

Displays your custom 405 page.

Force an internal server error (for example by loading a missing template) and verify that your custom 500 page appears.

---

## Why this matters

Large web applications often have dozens or hundreds of handlers.

Without reusable error handling, every handler contains duplicated code.

Centralizing this logic makes applications easier to maintain and is standard practice in production web servers.

---

## Think About

Write your answers as comments at the top of your helper file.

* Why shouldn't every handler parse and execute an error template separately?

* What is the difference between:

```go
http.Error(...)
```

and

```go
w.WriteHeader(...)
template.Execute(...)
```

* Why must `WriteHeader()` be called before writing the HTML?

* What happens if you call `WriteHeader()` after executing the template?

-------------------------------------------------------------------

# Exercise 2 — Logger Middleware

## Goal

Create middleware that logs every incoming request before passing it to the actual handler.

Example output:

```
[GET] /
[POST] /ascii-art
[GET] /ascii-art-switch
```

---

## Background

Currently, requests travel directly from the server to your handler.

```
Browser
    │
    ▼
Handler
```

Middleware inserts itself between them.

```
Browser
    │
    ▼
Middleware
    │
    ▼
Handler
```

The middleware performs extra work (logging, authentication, timing, etc.) before allowing the request to continue.

The handler should not know the middleware exists.

---

## Key Tasks

* Create a function with the signature:

```go
func Logger(next http.Handler) http.Handler
```

* Inside the middleware:

  * print the request method
  * print the requested path

* After logging, pass the request to the next handler.

* Wrap every route with the middleware.

---

## Success Criteria

Starting the server and visiting pages should produce logs like:

```
[GET] /
```

```
[POST] /ascii-art
```

```
[GET] /ascii-art-switch
```

The application should continue working exactly as before.

---

## Why this matters

Almost every production web server uses middleware.

Logging, authentication, compression, CORS, recovery, and rate limiting are almost always implemented this way.

Understanding middleware is one of the biggest steps toward understanding frameworks like Gin or Chi.

---

## Think About

* Why doesn't the logger need to know which handler comes next?

* What is the type `http.Handler`?

* Why does `Logger` return another `http.Handler`?

* How is wrapping handlers an example of the Decorator Pattern?

-------------------------------------------------------------------

# Exercise 3 — Serving Static Files

## Goal

Serve CSS and image files directly from the server instead of embedding styles inside HTML.

---

## Background

Browsers don't receive only HTML.

When an HTML page contains:

```html
<link rel="stylesheet" href="/static/css/style.css">
```

the browser automatically makes another request:

```
GET /static/css/style.css
```

Your server must know how to respond with the file's contents.

Go provides `http.FileServer` for this purpose.

---

## Key Tasks

* Create a directory structure:

```
static/
    css/
        style.css
    images/
```

* Move your inline CSS into `style.css`.

* Configure a file server rooted at the `static` directory.

* Mount it under the `/static/` URL path.

* Use `http.StripPrefix()` so the file server receives the correct file paths.

---

## Success Criteria

Opening your webpage should automatically load:

```
/static/css/style.css
```

Typing:

```
http://localhost:8080/static/css/style.css
```

into the browser should display the CSS file.

---

## Why this matters

Every web application serves static assets:

* CSS
* JavaScript
* Images
* Fonts
* Icons

Understanding how `FileServer` works removes the "magic" behind frameworks that serve static content automatically.

---

## Think About

* Why doesn't `FileServer` automatically know to remove `/static`?

* What would happen if you omitted `http.StripPrefix()`?

* Why does the URL path differ from the filesystem path?

* How does the browser know it should request the CSS file?

-------------------------------------------------------------------

# Exercise 4 — Request Timing Middleware

## Goal

Create middleware that measures how long each HTTP request takes to complete and logs the result to the terminal.

Example output:

```text
[GET] / 1.4ms
[POST] /ascii-art 3.2ms
[GET] /ascii-art-switch 0.9ms
```

The timer should start **before** the request reaches the handler and stop **after** the handler finishes processing.

---

## Background

When a browser sends a request, the server performs many operations before sending a response:

```text
Browser
    │
    ▼
Middleware
    │
    ▼
Handler
    │
    ▼
Template Execution
    │
    ▼
Response Sent
```

Since middleware sits *around* the handler, it can measure **everything** that happens during the request.

The flow becomes:

```text
Start Timer
     │
     ▼
Call Handler
     │
     ▼
Handler finishes
     │
     ▼
Stop Timer
     │
     ▼
Print elapsed time
```

Go's `time` package makes this simple:

```go
start := time.Now()

// ... some work ...

elapsed := time.Since(start)
```

The `time.Time` returned by `time.Now()` represents the current moment.

`time.Since(start)` calculates how much time has passed since that moment and returns a `time.Duration`.

---

## Key Tasks

* Create a middleware with the signature:

```go
func Timer(next http.Handler) http.Handler
```

* At the beginning of the middleware:

  * record the current time using `time.Now()`.

* Pass the request to the next handler.

* After the handler returns:

  * calculate the elapsed time using `time.Since()`.

* Log:

  * the HTTP method
  * the requested path
  * the total time taken

* Wrap all your existing routes with this middleware.

---

## Success Criteria

Running the server and making requests should produce logs similar to:

```text
[GET] / 800µs
```

```text
[POST] /ascii-art 2.7ms
```

```text
[GET] /ascii-art-switch 1.1ms
```

Your application should behave exactly as before—the only difference should be the additional timing information in the terminal.

---

## Why this matters

Performance is one of the most important aspects of backend development.

When an application becomes slow, one of the first questions developers ask is:

> **"How long does each request take?"**

Timing middleware provides that answer.

Almost every production web server includes some form of request timing, whether for:

* performance monitoring
* debugging slow endpoints
* collecting metrics
* identifying bottlenecks
* alerting when requests become too slow

Frameworks like Gin, Chi, Echo, and many cloud platforms include similar middleware out of the box.

Understanding how to build it yourself helps you understand how those frameworks work internally.

---

## Think About

Write your answers as comments at the top of your middleware file.

* Why is the timer started **before** calling the next handler instead of inside the handler?

* Why must the elapsed time be calculated **after** `next.ServeHTTP()` returns?

* What type does `time.Since()` return? How is it different from an `int` or a `float64`?

* What would happen if your handler never returned (for example, an infinite loop)?

* If you combine the Logger middleware from Exercise 2 with the Timer middleware, which should execute first? What differences would the order make?

* Why is middleware the ideal place for timing requests instead of adding timing code inside every handler?

---

## Challenge Extension (Optional)

Enhance your timer middleware to classify requests based on how long they take.

For example:

```text
[FAST] [GET] / 750µs
```

```text
[NORMAL] [POST] /ascii-art 8ms
```

```text
[SLOW] [GET] /history 125ms
```

Define your own thresholds (for example, under 5ms is "FAST", between 5ms and 50ms is "NORMAL", and above 50ms is "SLOW").

**Objective:** Learn how middleware can not only observe requests but also analyze and categorize them, a common technique used in monitoring and observability systems.

-------------------------------------------------------------------

# Exercise 5 — Download ASCII Result

## Goal

Allow users to download the generated ASCII art as a text file instead of only viewing it in the browser.

After generating ASCII art, a **"Download Result"** button should appear.

Clicking the button should send a request to:

```text
GET /download
```

The server should respond by prompting the browser to download a file named:

```text
ascii-art.txt
```

containing the generated ASCII art.

---

## Background

Up until now, every request your server has handled has returned content that the browser displays directly.

For example:

```text
Browser
    │
    ▼
GET /
    │
    ▼
Server returns HTML
    │
    ▼
Browser renders the page
```

However, browsers don't always display responses.

The **HTTP response headers** tell the browser how to handle the response.

Normally, the browser receives something like:

```http
Content-Type: text/html
```

which tells it:

> "Render this as a webpage."

But if the server includes a header like:

```http
Content-Disposition: attachment
```

the browser interprets the response differently:

> "Don't display this—download it as a file."

This means the **same data** can either be displayed in the browser or downloaded, depending on the response headers.

---

## Understanding `Content-Disposition`

The `Content-Disposition` header tells the browser what it should do with the response body.

Without it:

```text
Server
   │
   ▼
ASCII Art
   │
   ▼
Browser displays it
```

With it:

```text
Content-Disposition: attachment
```

the browser changes its behavior:

```text
Server
   │
   ▼
ASCII Art
   │
   ▼
Browser downloads it
```

You can also suggest a filename.

For example:

```http
Content-Disposition: attachment; filename="ascii-art.txt"
```

Most browsers will use that filename in the Save dialog.

---

## Key Tasks

### 1. Add a Download Button

After ASCII art has been generated, display a button or link such as:

```text
Download Result
```

The button should only appear when there is a result available.

---

### 2. Create a New Route

Register a new handler for:

```text
GET /download
```

This handler is responsible only for sending the downloadable file.

---

### 3. Make the ASCII Art Available

Your download handler needs access to the generated ASCII art.

Think about:

* Where is the ASCII art currently stored?
* Does the download handler have access to it?
* How can one handler make data available to another?

You are free to choose a simple solution for now.

---

### 4. Set the Appropriate Response Headers

Before writing the response body:

* Tell the browser the response is a downloadable attachment.
* Suggest the filename `ascii-art.txt`.
* Send the appropriate content type for plain text.

---

### 5. Write the ASCII Art to the Response

Instead of rendering a template, write the ASCII art directly to the response body.

The browser should receive the ASCII art exactly as it appears on the webpage.

---

## Success Criteria

1. Generate ASCII art.

2. A **Download Result** button appears.

3. Clicking it sends:

```text
GET /download
```

4. The browser opens its download dialog (or automatically downloads the file, depending on browser settings).

5. The downloaded file is named:

```text
ascii-art.txt
```

6. Opening the file in a text editor shows the generated ASCII art exactly as displayed on the webpage.

---

## Why this matters

Not every HTTP response is meant to be rendered in a browser.

Many web applications return downloadable content, such as:

* PDF invoices
* CSV reports
* ZIP archives
* Images
* Configuration files
* Log files

These are all sent using ordinary HTTP responses with different headers.

Understanding response headers is essential because they control **how the client interprets the response**, not just what data is sent.

---

## Think About

Write your answers as comments at the top of your download handler.

* Why does the browser display HTML but download a file when `Content-Disposition` is changed?

* What is the difference between a **response header** and a **response body**?

* Why must headers be set **before** writing the response body?

* Why doesn't the download handler need to execute an HTML template?

* Where should the generated ASCII art live so both the display handler and the download handler can access it?

* If two users generate different ASCII art at the same time, what could go wrong if you store the result in a single global variable?

---

## Challenge Extension (Optional)

Instead of always downloading:

```text
ascii-art.txt
```

allow the user to choose the filename.

For example:

```text
GET /download?filename=my-banner
```

The downloaded file should become:

```text
my-banner.txt
```

If no filename is provided, default to:

```text
ascii-art.txt
```

**Objective:** Learn how query parameters can influence not only the content of a response but also the metadata sent to the client through HTTP headers. This reinforces the relationship between routes, request data, response headers, and browser behavior.


-------------------------------------------------------------------


# Exercise 6 — ASCII Art History

## Goal

Allow users to revisit previously generated ASCII art.

Each time a user generates new ASCII art, the result should be saved in a history collection.

Create a new page:

```text
/history
```

that displays every previous generation in the order they were created.

Initially, the history only needs to exist while the server is running. Once the server stops, the history may be lost.

---

## Background

So far, your application behaves like a calculator:

```text
User submits input
       │
       ▼
Server generates result
       │
       ▼
Result is displayed
       │
       ▼
Request ends
```

Once another request is made, the previous result is forgotten.

Many real applications remember previous operations:

* Search history
* Recently viewed items
* Chat history
* Transaction history
* Order history

To do this, the server stores information somewhere before sending the response.

For this exercise, you'll store everything **in memory**.

```text
Request
   │
   ▼
Generate ASCII
   │
   ▼
Save to History
   │
   ▼
Display Result
```

Later, you'll learn how to store this data permanently using files and databases.

---

## Understanding In-Memory Storage

Your server is a running Go program.

As long as the program is running, variables remain in memory.

For example:

```text
Server Starts
      │
      ▼
History = []
```

First request:

```text
History = [
    Hello
]
```

Second request:

```text
History = [
    Hello
    World
]
```

Third request:

```text
History = [
    Hello
    World
    Go
]
```

When the server stops:

```text
History = []
```

Everything is lost because it only existed in RAM.

This is called **volatile storage**.

---

## Key Tasks

### 1. Create a History Collection

Create a variable capable of storing every generated ASCII result.

Each history entry should contain enough information to reproduce what the user generated.

Think about:

* the original text
* the banner used
* the generated ASCII art
* optionally, the time it was created

---

### 2. Save Every Successful Generation

Whenever ASCII art is generated successfully:

* create a new history entry
* add it to the history collection

Only successful generations should be saved.

Failed requests should not appear in the history.

---

### 3. Create a New Route

Register a new handler:

```text
GET /history
```

This handler should display all previously generated entries.

---

### 4. Create a History Page

Create a new template for the history page.

Each history entry should display:

* the original input text
* the banner that was used
* the generated ASCII art

Arrange them in a readable format.

---

### 5. Provide Navigation

Allow users to move between:

```text
/
```

and

```text
/history
```

without manually typing URLs.

---

## Success Criteria

Suppose the user generates:

```text
Hello
```

using the Standard banner.

Then:

```text
World
```

using the Shadow banner.

Finally:

```text
Go
```

using the Thinkertoy banner.

Visiting:

```text
/history
```

should display all three entries in the order they were created.

Refreshing the page should not remove the history.

Restarting the server may clear the history.

---

## Why this matters

Many web applications need to remember information between requests.

Examples include:

* Shopping carts
* Notifications
* Search history
* User activity
* Audit logs
* Chat conversations

This exercise introduces one of the most important ideas in backend development:

> **A request can produce data that the server keeps for future requests.**

Right now, that data lives in memory.

Later, you'll replace in-memory storage with permanent storage such as:

* JSON files
* SQLite
* PostgreSQL

The logic of your application will remain largely the same—the storage mechanism will simply change.

---

## Think About

Write your answers as comments at the top of your history handler.

* Why does the history survive multiple requests but disappear when the server stops?

* Why is a slice a suitable data structure for storing history?

* Why should failed requests not be added to the history?

* Should the newest entry appear first or last? Why?

* If multiple users access the server simultaneously, should they all share the same history? What problems could this create?

* If your history grows to 100,000 entries, what performance or memory concerns might arise?

---

## Challenge Extension 1 (Optional)

Add a timestamp to every history entry.

Each entry should show when it was generated.

For example:

```text
[14:32:08]

Input:
Hello

Banner:
Standard

Result:
...
```

**Objective:** Learn how servers record metadata alongside application data.

---

## Challenge Extension 2 (Optional)

Add a button beside every history entry:

```text
Regenerate
```

Clicking it should regenerate the same ASCII art without requiring the user to retype the input.

Think about:

* Which information must be stored to make regeneration possible?
* Should the regenerated result create a new history entry?

**Objective:** Understand how stored application state can be reused to reproduce previous actions.

---

## Challenge Extension 3 (Preparation for the Next Exercise)

Currently, every user shares the same history because it is stored globally in memory.

Ask yourself:

> **How could the server keep a different history for each user?**

Don't implement it yet—this question prepares you for future exercises on **sessions, cookies, and user-specific state**, where each visitor maintains their own independent history.


-------------------------------------------------------------------


# Exercise 7 — Upload Your Own Banner

## Goal

Allow users to upload their own ASCII banner file instead of choosing one from a predefined list.

Replace the banner `<select>` element with a file upload input.

When the form is submitted, the server should use the uploaded banner file to generate the ASCII art.

---

## Background

Until now, your application has only accepted **text-based form data**.

For example:

```text
Text: Hello
Banner: standard.txt
```

Both values are simple strings sent as part of the HTTP request.

A file upload is different.

Instead of sending only text, the browser sends both:

* text fields
* binary file data

This changes the request format from a normal form submission to a **multipart form**.

The request flow now becomes:

```text
Browser
   │
   ▼
Text Input
+
Banner File
   │
   ▼
multipart/form-data
   │
   ▼
Server
   │
   ▼
Read Uploaded File
   │
   ▼
Generate ASCII Art
```

Unlike previous exercises, the server must now extract **two different kinds of data** from the same request.

---

## Understanding Multipart Forms

A regular HTML form sends data like this:

```text
text=Hello
banner=standard
```

Everything is plain text.

When a file is included, the browser packages the request into multiple sections:

```text
Part 1:
Text = Hello

Part 2:
Banner File

Part 3:
Other form fields...
```

This format is called:

```text
multipart/form-data
```

Because the request is structured differently, the server must first parse it before accessing its contents.

Go provides:

```go
r.ParseMultipartForm(...)
```

to process the incoming multipart request.

After parsing, uploaded files can be accessed separately from normal form values.

---

## Key Tasks

### 1. Modify the HTML Form

Replace the banner selection dropdown:

```html
<select>
```

with a file upload field:

```html
<input type="file">
```

Ensure the form is configured to support file uploads.

---

### 2. Update the Request Handling

Since the request now contains files, parse it as a multipart form.

Think about:

* When should the form be parsed?
* What happens if parsing fails?

Handle any parsing errors appropriately.

---

### 3. Read the Uploaded Banner

Retrieve the uploaded file from the request.

Think about:

* How do you know whether the user actually uploaded a file?
* What should happen if no file is provided?

Once you obtain the file, remember to release any resources associated with it when you're finished.

---

### 4. Use the Uploaded Banner

Your current ASCII generator expects a banner stored on disk.

Modify your application so it can use the uploaded banner instead.

Think carefully about the design.

Possible questions include:

* Should the generator read from a filename?
* Should it read from file contents?
* Should it accept any data source?

Choose an approach that makes your generator more flexible.

---

### 5. Validate the Banner

Before generating ASCII art, verify that the uploaded banner is valid.

Consider questions such as:

* Is the file empty?
* Does it contain the expected number of lines?
* Does it appear to be a valid ASCII banner?

Reject invalid banner files with an appropriate error.

---

## Success Criteria

A user should be able to:

1. Enter text.

2. Select a banner file from their computer.

3. Submit the form.

4. Receive ASCII art generated using the uploaded banner.

If:

* no banner is uploaded,
* the file cannot be read,
* or the banner is invalid,

the server should return an appropriate error instead of crashing.

---

## Why this matters

File uploads are one of the most common features in web applications.

Examples include:

* Profile pictures
* PDF documents
* Videos
* Spreadsheets
* Configuration files
* CSV imports

Although the uploaded content changes, the overall workflow is the same:

```text
Browser
    │
    ▼
Upload File
    │
    ▼
Server Validates
    │
    ▼
Server Processes
    │
    ▼
Response Sent
```

Understanding multipart forms is an essential backend skill because many real-world applications depend on receiving files from users.

---

## Think About

Write your answers as comments at the top of your upload handler.

* Why can't uploaded files be read using `FormValue()`?

* What is the purpose of `ParseMultipartForm()`?

* Why must the HTML form be configured differently when uploading files?

* Why is it important to validate uploaded files before processing them?

* Why should uploaded files be closed after you're finished reading them?

* Your current generator reads banner files from disk. Is that the best design now that banners can come from an uploaded file? How could you redesign the generator to work with both built-in and uploaded banners?

---

## Challenge Extension 1 (Optional)

Support **both** methods of selecting a banner:

* Choose one of the built-in banners.
* Upload a custom banner.

If both are provided, decide which one takes priority.

Document your decision.

**Objective:** Learn how to design application logic that supports multiple input sources cleanly.

---

## Challenge Extension 2 (Optional)

Display information about the uploaded banner before generating the ASCII art.

For example:

```text
Filename:
mybanner.txt

Size:
6.2 KB

Upload Status:
Valid
```

**Objective:** Learn that uploaded files come with metadata, not just file contents.

---

## Challenge Extension 3 (Preparation for Future Exercises)

Currently, every request uploads the banner again, even if it's the same file.

Ask yourself:

> **How could the server save uploaded banners and allow users to reuse them later without uploading them every time?**

Don't implement this yet. This question introduces ideas you'll encounter later when working with persistent storage, databases, and file management systems.