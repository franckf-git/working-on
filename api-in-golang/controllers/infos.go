package controllers

import (
	"encoding/json"
	"lite-api-crud/config"
	"net/http"
)

func WelcomePage(res http.ResponseWriter, req *http.Request) {
	homepage := config.Message{
		Status:  "information",
		Message: config.WelcomeMessage,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(homepage)
}

func Docs(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, config.DocsLink, http.StatusMovedPermanently)
}

func NotFoundMessage(res http.ResponseWriter, req *http.Request) {
	notfound := config.Message{
		Status:  "error",
		Message: "this route doesn't exist",
	}
	res.WriteHeader(http.StatusNotFound)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(notfound)
}
