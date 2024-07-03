package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl = template.Must(template.New("").ParseFiles("views.gotmpl"))

	http.HandleFunc("/", urlParam)
	http.HandleFunc("/form", formParam)
	http.HandleFunc("/uploadForm", formUpload)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func urlParam(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Hello!, "+v)
}

func formParam(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "formOut", req.FormValue("q"))
}

func formUpload(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "formUpload", s)
}
