Week 1 — HTML + Your First Go Server
Days 1–2 · HTML (3–4 hrs total)

You only need to learn the subset of HTML relevant to this project. Do not do a full HTML course.

Go here: https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Getting_started

Read and do the exercises for only these pages in order:

    Getting started with HTML — understand tags, elements, attributes
    The head — skip most of it, just know <title>
    HTML text fundamentals — headings, paragraphs
    Creating hyperlinks — skip
    Web forms → Your first form — https://developer.mozilla.org/en-US/docs/Learn/Forms/Your_first_form — this is the most important page
    Web forms → Basic native form controls — https://developer.mozilla.org/en-US/docs/Learn/Forms/Basic_native_form_controls — focus on text input, radio buttons, select, and submit button

By end of Day 2 you should be able to write this by hand without looking:
html

<!DOCTYPE html>
<html>
  <head><title>ASCII Art</title></head>
  <body>
    <form action="/ascii-art" method="POST">
      <input type="text" name="text" />
      <input type="radio" name="banner" value="standard" /> Standard
      <input type="radio" name="banner" value="shadow" /> Shadow
      <input type="radio" name="banner" value="thinkertoy" /> Thinkertoy
      <button type="submit">Generate</button>
    </form>
  </body>
</html>

Days 3–4 · Go net/http (4–6 hrs total)

Go here: https://gobyexample.com

Read these pages in this exact order:

    HTTP Servers — https://gobyexample.com/http-servers
    HTTP Server — writing handlers — same page, read carefully
    Then read the official docs page for the two functions you will actually use:
        http.HandleFunc — https://pkg.go.dev/net/http#HandleFunc
        http.ListenAndServe — https://pkg.go.dev/net/http#ListenAndServe

What to build on Day 4: Write a Go file from scratch that:

    Starts a server on port 8080
    Has a GET / handler that just writes "Hello" to the browser
    Has a POST /ascii-art handler that reads r.FormValue("text") and prints it to the terminal

This is your proof that you understand it before moving on. Do not skip this.

Key things to know cold:
go

// Reading form values from a POST request
text := r.FormValue("text")
banner := r.FormValue("banner")

// Sending HTTP error codes
http.Error(w, "Not found", http.StatusNotFound)       // 404
http.Error(w, "Bad request", http.StatusBadRequest)   // 400
http.Error(w, "Server error", http.StatusInternalServerError) // 500

Days 5–6 · Go html/template (4–6 hrs total)

This is where Go sends data to your HTML page.

Read this, the whole page: https://pkg.go.dev/html/template

Then read this short practical tutorial which is clearer: https://gobyexample.com/text-templates (text templates) then https://pkg.go.dev/html/template#example-package (the html/template example at the bottom)

The 4 things to learn from this:

    How to parse a template file:

go

tmpl, err := template.ParseFiles("templates/index.html")

    How to execute it (send it to the browser):

go

tmpl.Execute(w, data)

    How to pass a Go struct as data:

go

type PageData struct {
    Result string
}
data := PageData{Result: "your ascii art here"}
tmpl.Execute(w, data)

    How to display that data inside your HTML file:

html

<pre>{{.Result}}</pre>

What to build on Day 6: Combine Days 4–6. Make a page that accepts text input, and displays it back on the page in a <pre> tag. Not ascii art yet — just echo the text back. If this works, you have learned everything structural you need.
Week 2 — Connect Everything + Finish
Day 7 · Wire in your ascii-art logic

    Move your ascii-art code into a function: func GenerateAscii(text, banner string) (string, error)
    Call it from your POST handler
    Pass the result to your template
    Display it in <pre>{{.Result}}</pre>

Day 8 · HTTP status codes

Handle every error case properly:

    Template file missing → 500
    Banner file missing → 404
    Empty text input → 400
    Everything OK → 200 (default, nothing needed)

Go here for reference: https://pkg.go.dev/net/http#pkg-constants — search for StatusOK, StatusBadRequest etc.
Day 9 · Project structure

Your folder must look exactly like this:

ascii-art-web/
├── main.go
├── templates/
│   └── index.html
├── shadow.txt
├── standard.txt
├── thinkertoy.txt
└── README.md

Days 10–11 · Testing

Test every case:

    Normal input with each of the 3 banners
    Empty text box submitted
    Manually visit a bad URL like /badroute — should get 404
    Put a wrong template filename in code temporarily — should get 500

Days 12–13 · README + code cleanup

Write README.md with these 4 sections the project requires: Description, Authors, Usage, Implementation details. Keep it short and honest.

Review your code: every err must be checked, no variable called x or temp, functions should do one thing.
Day 14 · Buffer day

Finish anything incomplete. Re-read the project spec line by line and check every requirement is met.
The 3 URLs You will Return To Most
What	URL
HTML forms reference	https://developer.mozilla.org/en-US/docs/Learn/Forms/Basic_native_form_controls
net/http package	https://pkg.go.dev/net/http
html/template package	https://pkg.go.dev/html/template


Days 1–2 · HTML

Day 1 — HTML basics Traversy Media — HTML Crash Course For Absolute Beginners https://www.youtube.com/watch?v=UB1O30fR-EE ~1 hour. Covers all the tags and structure you need for Day 1.

Day 2 — HTML Forms There is no standalone Traversy Media forms-only video. Use this instead — it is clearer and more focused on exactly what you need: Web Dev Simplified — HTML Forms https://www.youtube.com/watch?v=fNcJuPIZ2WE ~15 minutes. Covers text input, radio buttons, select, and submit — exactly your Day 2 tasks.
Days 3–4 · Go net/http

Go Web Server from scratch Anthony GG — Building a Web Server in Go https://www.youtube.com/watch?v=FmghoRV3E6s Covers HandleFunc, ListenAndServe, reading POST form values — maps directly to Days 3 and 4.
Days 5–6 · Go html/template

Go Templates Golang Dojo — Go Templates Tutorial https://www.youtube.com/watch?v=k5MoTn6FBjQ Covers ParseFiles, Execute, passing structs, and displaying data in HTML — your exact Days 5 and 6.
Days 8–9 · Full Go Web App (wiring it all together)

Complete Go web app walkthrough TutorialEdge — Creating a Simple Web Server with Golang https://www.youtube.com/watch?v=iJnBuoqkdLM Shows the full flow — form to handler to template to browser — which is exactly what you are building on Days 8 and 9.
Day 10 · HTTP Status Codes

Fireship — Every HTTP Response Code Explained https://www.youtube.com/watch?v=wJa5CTIFj7U Under 10 minutes. Covers 200, 400, 404, 500 in a memorable way.
The One Rule When Using These Videos

Pause every 2–3 minutes and type what you just watched. Do not watch a full video and then try to code — you will forget 70% of it. Watch a small chunk, pause, type it, then continue.
Attachments area
Preview YouTube video HTML Crash Course For Absolute BeginnersPreview YouTube video HTML Crash Course For Absolute Beginners
Preview YouTube video Learn HTML Forms In 25 MinutesPreview YouTube video Learn HTML Forms In 25 Minutes
Preview YouTube video Learn HTTP Status Codes In 10 MinutesPreview YouTube video Learn HTTP Status Codes In 10 Minutes
