package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shinshin86/mini-speech-bubbler/static"
)

var tpl *template.Template

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imagePath := "/static/avatar.png"
	htmlPath := "index.html"

	f, err := static.Root.Open(htmlPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New(htmlPath).Parse(string(b))
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(w, imagePath)
	if err != nil {
		panic(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", viewHandler)
	log.Print("Access to localhost:3000")
	http.ListenAndServe(":3000", nil)
}
