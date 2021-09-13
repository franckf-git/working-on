package main

import (
	"log"
	"net/http"

	"lite-api-crud/config"
	"lite-api-crud/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.WelcomePage)
	router.HandleFunc("/api/v1/posts", controllers.ShowAllPosts)
	log.Fatal(http.ListenAndServe(config.PORT, router))
}
