package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	indextemplate, err := template.ParseFiles("indextemplate.html")
	if err != nil {
		panic(err)
	}

	title := "Titre de la page de template"
	subtitle := "Sous-titre du héros"
	data := struct {
		Title    string
		SubTitle string
	}{
		Title:    title,
		SubTitle: subtitle,
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		indextemplate.Execute(res, data)
	})
	log.Print(http.ListenAndServe(":5500", nil))

}
