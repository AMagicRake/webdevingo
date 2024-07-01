package main

import (
	"io"
	"net/http"
)

func ServeDog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}

func ServeCat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}
func ServeIndex(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello!")
}

func main() {
	// mux := http.NewServeMux()
	http.HandleFunc("/", ServeIndex)
	http.HandleFunc("/dog", ServeDog)
	http.HandleFunc("/cat", ServeCat)
	http.ListenAndServe(":8080", nil)
}
