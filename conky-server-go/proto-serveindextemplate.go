package main

import (
	"html/template"
	"os"
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

	err = indextemplate.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
