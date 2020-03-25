package main

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintf(w, "Executed /a\n")
	case "/b":
		fmt.Fprintf(w, "Executed /b\n")
	case "/c":
		fmt.Fprintf(w, "Executed /c\n")
	default:
		http.Error(w, "404 Not Found\n", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}
