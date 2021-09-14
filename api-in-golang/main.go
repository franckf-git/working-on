package main

import (
	"log"
	"net/http"

	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"lite-api-crud/models"

	"github.com/gorilla/mux"
)

func init() {
	db := models.OpenDatabase()
	defer db.Close()
	models.StartDatabase(db)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.WelcomePage).Methods("GET")
	router.HandleFunc("/api/v1/posts", controllers.ShowAllPosts).Methods("GET")
	router.HandleFunc("/api/v1/post", controllers.AddPost).Methods("POST")
	log.Println("api server is up")
	log.Fatal(http.ListenAndServe(config.PORT, router))
}
