package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func main() {
	tpl = template.Must(template.ParseFiles("views.gotmpl"))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("sessionID")
	if err != nil {
		sessionId := uuid.New()
		newCookie := &http.Cookie{Name: "sessionID", Value: sessionId.String()}
		fmt.Println(newCookie)
		http.SetCookie(w, newCookie)
	}

	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index", u)

}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar", u)
}
