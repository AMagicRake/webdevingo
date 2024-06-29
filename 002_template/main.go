package main

import (
	"log"
	"os"
	"text/template"
)

type data struct {
	Heading string
}

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index2", data{"Test Replace"})
	if err != nil {
		log.Fatalln(err)
	}
}
