package controllers

import (
	"encoding/json"
	"lite-api-crud/config"
	"net/http"
)

func WelcomePage(res http.ResponseWriter, req *http.Request) {
	homepage := struct {
		Message string `json:"message"`
	}{
		Message: config.WelcomeMessage,
	}
	json.NewEncoder(res).Encode(homepage)
}

func Docs(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "https://gitlab.com/franckf/working-on/-/blob/master/api-in-golang/readme.md#documentation-de-lapi", http.StatusMovedPermanently)
}
