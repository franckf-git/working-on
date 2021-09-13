package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"lite-api-crud/config"

	"github.com/gorilla/mux"
)

func WelcomePage(res http.ResponseWriter, req *http.Request) {
	homepage := struct {
		Message string `json:"message"`
	}{
		Message: config.WelcomeMessage,
	}
	json.NewEncoder(res).Encode(homepage)
}

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "all posts")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", WelcomePage)
	router.HandleFunc("/api/v1/posts", ShowAllPosts)
	log.Fatal(http.ListenAndServe(config.PORT, router))
}
