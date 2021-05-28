package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

// lister toutes commandes à retourner dans la page web
var command map[string]string = map[string]string{
	"uptime":   "uptime",
	"user":     "whoami",
	"ips":      "ip a",
	"hostname": "hostname",
	"packages": "dnf list --installed | wc -l", // prévoir pour debian distro
	"kernel":   "uname -a",
	// "os":       "lsb_release -a",
	"top":    "ps aux | sort -nk +4 | tail",
	"memory": "free -h",
	"loads":  "uptime | cut -d' ' -f10-",
	"cpu":    "lscpu | grep 'Model name' | cut -d' ' -f12-",
	"disks":  "df -h",
}

func main() {
	// dirty

	for _, v := range command {
		// les executer
		// conncurrence ? await ? channels ?
		log.Println(runCommand(v))
	}

	// servir le template
	http.HandleFunc("/", serveTemplate)
	log.Println("Le serveur est en ligne, visitez http://127.0.0.1:5500")
	http.ListenAndServe(":5500", nil)

	// ouverture auto du navigateur par défaut ? `xdg-open http://127.0.0.1:5500`

	//! **refacto**

	return
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	// recupérer le template
	indextemplate, err := template.ParseFiles("indextemplate.html")
	if err != nil {
		panic(err)
	}

	// definir les champs du template avec les infos des commandes externes
	// prévoir une boucle dans le template
	title := "ConkyWeb"
	data := struct {
		Title    string
		Commands map[string]string
	}{
		Title: title,
		Commands: map[string]string{
			"1": "a",
			"2": "b",
			"3": "c",
			"9": "a",
			"8": "b",
			"7": "c",
			"6": "a",
			"5": "b",
			"4": "c",
		},
	}
	indextemplate.Execute(res, data)
}

func runCommand(command string) string {
	output, errcmd := exec.Command("bash", "-c", command).CombinedOutput()
	if errcmd != nil {
		log.Fatal("La commande ", command, " n'existe pas")
	}
	result := string(output)
	return result
}
