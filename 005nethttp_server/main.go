package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//implement hotdog?
	fmt.Println(r.RequestURI, r.Method)
	fmt.Fprintln(w, "You wanted a hotdog?")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
