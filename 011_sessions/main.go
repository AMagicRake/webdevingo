package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var (
	tpl               *template.Template
	dbUsers           = map[string]user{}
	dbSessions        = map[string]session{}
	dbSessionsCleaned time.Time
)

const sessionLength time.Duration = (time.Second * 30)

func init() {
	tpl = template.Must(template.ParseFiles("views.gotmpl"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("sessionID")
	if err != nil {
		sessionId := uuid.New()
		c := &http.Cookie{
			Name:   "sessionID",
			Value:  sessionId.String(),
			MaxAge: int(sessionLength / time.Second),
		}
		fmt.Println(c)
		http.SetCookie(w, c)
	}

	var u user
	if session, ok := dbSessions[c.Value]; ok {
		u = dbUsers[session.un]
	}

	if req.Method == http.MethodPost {
		uName := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{uName, f, l}
		dbSessions[c.Value] = session{uName, time.Now()}
		dbUsers[uName] = u
	}

	if time.Since(dbSessionsCleaned) > sessionLength {
		for k, v := range dbSessions {
			if time.Since(v.lastActivity) > sessionLength {
				delete(dbSessions, k)
			}
		}
		dbSessionsCleaned = time.Now()
	}

	tpl.ExecuteTemplate(w, "index", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("sessionID")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	session, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	session.lastActivity = time.Now()
	dbSessions[c.Value] = session
	u := dbUsers[session.un]
	c.MaxAge = int(sessionLength / time.Second)
	// update cookie length
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "bar", u)
}
