package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shinshin86/mini-speech-bubbler/static"
)

var tpl *template.Template

func avatarImagePath() (string, error) {
	pngPath := filepath.FromSlash("./static/avatar.png")
	if _, err := os.Stat(pngPath); err == nil {
		return pngPath, nil
	}

	jpgPath := filepath.FromSlash("./static/avatar.jpg")
	if _, err := os.Stat(jpgPath); err == nil {
		return jpgPath, nil
	}

	jpegPath := filepath.FromSlash("./static/avatar.jpeg")
	if _, err := os.Stat(jpegPath); err == nil {
		return jpegPath, nil
	}

	return "", errors.New("No avatar image is set")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imagePath, err := avatarImagePath()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

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
