package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/shocked_pikachu.jpg", shockedPikachu)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./img"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/img/shocked_pikachu.jpg">`)
}

func shockedPikachu(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("img/shocked_pikachu.jpg")
	if err != nil {
		http.Error(w, "File not Found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not Found", 404)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
