package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//implement hotdog?
	fmt.Println(r.Method, r.RequestURI, r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	json, err := json.MarshalIndent(r.Form, "", "  ")
	fmt.Println(string(json))

	tpl.ExecuteTemplate(w, "index.gotmpl", r.Form)
}

var tpl *template.Template

func main() {
	var d hotdog
	tpl = template.Must(template.ParseFiles("index.gotmpl"))
	http.ListenAndServe(":8080", d)
}
