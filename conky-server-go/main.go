package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// dirty

	// lister toutes commandes à retourner dans la page web

	// les executer
	// conncurrence ? await ? channels ?
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error : %v", err)
	}
	log.Println(hostname)

	output, errcmd := exec.Command("uptime").CombinedOutput()
	if errcmd != nil {
		os.Stderr.WriteString(errcmd.Error())
	}
	result := string(output)
	log.Print(result)

	// recupérer le template
	indextemplate, err := template.ParseFiles("indextemplate.html")
	if err != nil {
		panic(err)
	}

	// definir les champs du template avec les infos des commandes externes
	// prévoir une boucle dans le template

	title := "Titre de la page de template"
	subtitle := "Sous-titre du héros"
	data := struct {
		Title    string
		SubTitle string
	}{
		Title:    title,
		SubTitle: subtitle,
	}

	// servir le template
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		indextemplate.Execute(res, data)
	})
	log.Print(http.ListenAndServe(":5500", nil))

	// ouverture auto du navigateur par défaut ?

	//! **refacto**

	return
}
