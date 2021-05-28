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

	var command map[string]string = map[string]string{
		"uptime":   "uptime",
		"user":     "whoami",
		"ips":      "ip a",
		"hostname": "hostname",
		"packages": "dnf list --installed | wc -l", // prévoir pour debian distro
		"kernel":   "uname -a",
		"os":       "lsb_release -a", // prévoir une gestion des erreurs si la commande n'est disponible
		"top":      "ps aux | sort -nk +4 | tail",
		"memory":   "free -h",
		"loads":    "uptime | cut -d' ' -f10-",
		"cpu":      "lscpu | grep 'Model name' | cut -d' ' -f12-",
		"disks":    "df -h",
	}

	for k, v := range command {
		// les executer
		// conncurrence ? await ? channels ?
		log.Println("the key is", k, "- the value is", v)
		log.Println(runCommand(v))
	}

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

	// ouverture auto du navigateur par défaut ? `xdg-open http://127.0.0.1:5500`

	//! **refacto**

	return
}

func runCommand(command string) string {
	output, errcmd := exec.Command("bash", "-c", command).CombinedOutput()
	if errcmd != nil {
		os.Stderr.WriteString(errcmd.Error())
	}
	result := string(output)
	return result
}
