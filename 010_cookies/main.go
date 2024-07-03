package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexPage(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		fmt.Println("No Cookie Found")
	}
	if err != nil {
		newCookie := &http.Cookie{Name: "my-cookie", Value: "Niel"}
		fmt.Println(newCookie)
		http.SetCookie(w, newCookie)
		io.WriteString(w, "Hi! Nice to meet you!\n")
	} else {
		io.WriteString(w, "Welcome Back, "+c.Value+"!\n")
	}
	io.WriteString(w, "I am the cookie monster!\n")

}
