package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func startServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/greetings", greetings)

	http.ListenAndServe(":8008", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func greetings(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "greetings\n")
}
