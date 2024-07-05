package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func main() {

	http.HandleFunc("/", indexPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, req *http.Request) {

	sessionId := uuid.New()
	c, err := req.Cookie("sessionID")
	if err != nil {
		newCookie := &http.Cookie{Name: "sessionID", Value: sessionId.String()}
		fmt.Println(newCookie)
		http.SetCookie(w, newCookie)
		io.WriteString(w, "Hi! Nice to meet you!\n")
	} else {
		io.WriteString(w, "Welcome Back, "+c.Value+"!\n")
	}
	io.WriteString(w, "I am the cookie monster!\n")

}
