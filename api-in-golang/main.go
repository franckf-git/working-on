package main

import (
	"log"
	"net/http"

	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"lite-api-crud/models"

	"github.com/gorilla/mux"
)

func main() {
	db := models.OpenDatabase()
	defer db.Close()
	models.StartDatabase(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.WelcomePage)
	router.HandleFunc("/api/v1/posts", controllers.ShowAllPosts)
	log.Fatal(http.ListenAndServe(config.PORT, router))
}
