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

	// tpl, err := template.ParseFiles("tpl.gotmpl")
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gotmpl"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	tpl.ExecuteTemplate(os.Stdout, "index2", data{"Test Replace"})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
