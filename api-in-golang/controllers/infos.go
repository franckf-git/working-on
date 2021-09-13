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
