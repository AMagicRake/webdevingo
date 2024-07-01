package main

import (
	"io"
	"net/http"
)

type dog int
type cat int
type index int

func (dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}

func (cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}
func (index) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello!")
}

func main() {
	var d dog
	var c cat
	var i index
	mux := http.NewServeMux()
	mux.Handle("/", i)
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}
