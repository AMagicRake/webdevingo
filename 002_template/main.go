package main

import (
	"os"
	"text/template"
)

type data struct {
	Heading string
}

func helloWorld() string {
	return "Hello World"
}

func newDiv(data string) string {
	return "<div>" + data + "</div>"
}

func main() {

	fm := template.FuncMap{
		"hw": helloWorld,
		"nd": newDiv,
	}

	tpl := template.New("")
	tpl = tpl.Funcs(fm)
	tpl = template.Must(tpl.ParseFiles("tpl.gotmpl"))

	tpl.ExecuteTemplate(os.Stdout, "index2", data{"Test Replace"})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
