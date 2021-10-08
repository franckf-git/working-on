package router

import (
	"log"
	"net/http"

	"lite-api-crud/config"
	"lite-api-crud/controllers"

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
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.Use(setHeader)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controllers.WelcomePage).Methods("GET")
	a.Router.HandleFunc("/api/v1/docs", controllers.Docs).Methods("GET")
	a.Router.HandleFunc("/api/v1/posts", controllers.ShowAllPosts).Methods("GET")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", controllers.ShowPost).Methods("GET")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", isAuthorized(controllers.UpdatePost)).Methods("PUT")
	a.Router.HandleFunc("/api/v1/post/{id:[0-9]+}", isAuthorized(controllers.DeletePost)).Methods("DELETE")
	a.Router.HandleFunc("/api/v1/post", isAuthorized(controllers.AddPost)).Methods("POST")
	a.Router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	a.Router.HandleFunc("/user/jwt", controllers.AskJWT).Methods("POST")
	a.Router.NotFoundHandler = http.HandlerFunc(controllers.NotFoundMessage)
}
