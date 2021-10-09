package router

import (
	"log"
	"net/http"
	"time"

	"lite-api-crud/config"
	"lite-api-crud/controllers"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run() {
	log.Println("ENV:", config.State)
	server := &http.Server{
		Addr:              config.PORT,
		Handler:           a.Router,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 60,
		MaxHeaderBytes:    0,
		ErrorLog:          &log.Logger{},
	}
	log.Println("api server is up")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.Use(setHeader)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controllers.WelcomePage).Methods("GET")
	api := a.Router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	postReading := v1.PathPrefix("/post").Subrouter()
	post := v1.PathPrefix("/post").Subrouter()
	user := a.Router.PathPrefix("/user").Subrouter()

	v1.HandleFunc("/docs", controllers.Docs).Methods("GET")
	v1.HandleFunc("/posts", controllers.ShowAllPosts).Methods("GET")

	postReading.Use(checkContentType)
	postReading.HandleFunc("", isAuthorized(controllers.AddPost)).Methods("POST")
	postReading.HandleFunc("/{id:[0-9]+}", isAuthorized(controllers.UpdatePost)).Methods("PUT")
	post.HandleFunc("/{id:[0-9]+}", controllers.ShowPost).Methods("GET")
	post.HandleFunc("/{id:[0-9]+}", isAuthorized(controllers.DeletePost)).Methods("DELETE")

	user.Use(checkContentType)
	user.HandleFunc("", controllers.AddUser).Methods("POST")
	user.HandleFunc("/jwt", controllers.AskJWT).Methods("POST")

	a.Router.NotFoundHandler = http.HandlerFunc(controllers.NotFoundMessage)
}
