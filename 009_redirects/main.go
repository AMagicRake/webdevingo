package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tmp *template.Template

func main() {

	tmp = template.Must(template.New("").ParseFiles("views.gotmpl"))

	http.HandleFunc("/", index)
	http.HandleFunc("/seeOther", seeOther)
	http.HandleFunc("/testForm", testForm)
	http.HandleFunc("/temporaryRedirect", temporaryRedirect)
	http.HandleFunc("/movedPermenantly", movedPermenantly)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method in index:", req.Method)
	io.WriteString(w, "Your request method at index: "+req.Method)
}

func testForm(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method in testForm:", req.Method)
	tmp.ExecuteTemplate(w, "index", nil)
}

func seeOther(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method in seeOther:", req.Method)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func temporaryRedirect(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method in temporaryRedirect:", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func movedPermenantly(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Method in movedPermenantly:", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
