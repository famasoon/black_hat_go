package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("star")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":8080", &l)
}
