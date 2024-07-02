package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/shocked_pikachu.jpg", shockedPikachu)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/shocked_pikachu.jpg">`)
}

func shockedPikachu(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("shocked_pikachu.jpg")
	if err != nil {
		http.Error(w, "File not Found", 404)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not Found", 404)
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
