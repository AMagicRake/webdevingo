package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", urlParam)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func urlParam(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Hello!, "+v)
}
