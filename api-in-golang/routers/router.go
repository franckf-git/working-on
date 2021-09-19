package router

import (
	"log"
	"net/http"

	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"lite-api-crud/models"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run() {
	log.Println("api server is up")
	log.Fatal(http.ListenAndServe(config.PORT, a.Router))
}

func (a *App) Initialize() {
	models.InitializeDB()
	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controllers.WelcomePage).Methods("GET")
	a.Router.HandleFunc("/api/v1/docs", controllers.Docs).Methods("GET")
	a.Router.HandleFunc("/api/v1/posts", controllers.ShowAllPosts).Methods("GET")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", controllers.ShowPost).Methods("GET")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", controllers.UpdatePost).Methods("PUT")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", controllers.DeletePost).Methods("DELETE")
	a.Router.HandleFunc("/api/v1/post", controllers.AddPost).Methods("POST")
}
