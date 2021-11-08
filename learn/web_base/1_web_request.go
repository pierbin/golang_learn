package main

import (
	"fmt"
	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 9000
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	http.ListenAndServe(":9000", nil)
}

/*
	First, create a Handler which receives all coming HTTP connections from browsers, HTTP clients or API requests. A handler in Go is a function with this signature:
	func (w http.ResponseWriter, r *http.Request)

	An http.ResponseWriter which is where you write your text/html response to.
	An http.Request which contains all information about this HTTP request including things like the URL or header fields.

	Second, Registering a request handler to the default HTTP Server is as simple as this: "http.HandleFunc"
*/
func handler(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path will output the url name after '/', for example: http://localhost:9000/test, it will output "/test"
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)  //r.URL.Path[0:] is the same with r.URL.Path.
	fmt.Fprintf(w, "Hello, your route is: %s\n", r.URL.Path[1:]) //it will output "test"
	fmt.Fprint(w, "Welcome to my website")

	//http.FileServer and point it to a url path. For the file server to work properly it needs to know, where to serve files from, using http.FileServer to serve static assets like JavaScript, CSS and images
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
}
